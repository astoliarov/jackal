package app

import (
	"bytes"
	"github.com/oliamb/cutter"
	"image"

	"image/gif"
	"image/jpeg"
	"image/png"
)

type CropType int

const CropTypeDefault CropType = 1
const CropTypeRatio CropType = 2

type ImageProcessingService struct{}

func (service *ImageProcessingService) CropCentered(content []byte, width int, height int, cropType CropType) ([]byte, string, error) {
	reader := bytes.NewReader(content)

	img, imageType, err := image.Decode(reader)
	if err != nil {
		return nil, "", err
	}

	config := cutter.Config{
		Width:  width,
		Height: height,
		Mode:   cutter.Centered,
	}

	if cropType == CropTypeRatio {
		config.Options = cutter.Ratio
	}

	croppedImg, err := cutter.Crop(img, config)

	buffer := new(bytes.Buffer)

	switch imageType {
	case "jpeg":
		jpeg.Encode(buffer, croppedImg, nil)
	case "png":
		png.Encode(buffer, croppedImg)
	case "gif":
		gif.Encode(buffer, croppedImg, nil)
	}

	return buffer.Bytes(), imageType, nil
}

func NewImageProcessingService() *ImageProcessingService {
	service := ImageProcessingService{}
	return &service
}
