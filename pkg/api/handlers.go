package api

import (
	"github.com/gin-gonic/gin"
	"jackal/pkg/interfaces"
	"net/http"
	"strconv"
)

type CropRequest struct {
	Url    string
	Height int
	Width  int
	Type   interfaces.CropType
}

type CropHandler struct {
	useCase interfaces.IDownloadAndCropImageUseCase
}

func (h *CropHandler) fillAndValidate(c *gin.Context, req *CropRequest) ValidateErrs {
	validateErrs := make(ValidateErrs)

	url := c.Query("url")
	if url == "" {
		validateErrs["url"] = "field is required"
	} else {
		req.Url = url
	}

	widthRaw := c.Query("width")
	if widthRaw == "" {
		validateErrs["width"] = "field is required"
	} else {

		value, err := strconv.Atoi(widthRaw)
		if err != nil {
			validateErrs["width"] = "bad value"
		} else {
			req.Width = value
		}

	}

	heightRaw := c.Query("height")
	if heightRaw == "" {
		validateErrs["height"] = "field is required"
	} else {
		value, err := strconv.Atoi(heightRaw)
		if err != nil {
			validateErrs["height"] = "bad value"
		} else {
			req.Height = value
		}
	}

	cropType := c.Query("type")
	if cropType == "ratio" {
		req.Type = interfaces.CropTypeRatio
	} else {
		req.Type = interfaces.CropTypeDefault
	}

	return validateErrs
}

func (h *CropHandler) HandleGetCrop(c *gin.Context) {

	var cropRequest CropRequest

	errs := h.fillAndValidate(c, &cropRequest)
	if !errs.IsEmpty() {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	var croppedImage []byte
	var imageType string
	var err error

	croppedImage, imageType, err = h.useCase.Execute(cropRequest.Url, cropRequest.Width, cropRequest.Height, cropRequest.Type)

	if err != nil {
		c.Data(http.StatusNoContent, "application/empty", []byte{})
		return
	}

	var contentType string
	switch imageType {
	case "jpeg":
		contentType = "image/jpeg"
	case "png":
		contentType = "image/png"
	case "gif":
		contentType = "image/gif"
	}
	c.Data(http.StatusOK, contentType, croppedImage)
}

func NewCropHandler(cropUseCase interfaces.IDownloadAndCropImageUseCase) *CropHandler {
	service := CropHandler{}
	service.useCase = cropUseCase
	return &service
}
