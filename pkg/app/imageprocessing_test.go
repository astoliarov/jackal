package app

import (
	"testing"
	"github.com/stretchr/testify/suite"
	"image"
	"image/color"
	"bytes"
	"image/jpeg"
	"encoding/base64"
	"image/png"
	"image/gif"
)

const croppedImg = `/9j/2wCEAAgGBgcGBQgHBwcJCQgKDBQNDAsLDBkSEw8UHRofHh0aHBwgJC4nICIsIxwcKDcpLDAxNDQ0Hyc5PTgyPC4zNDIBCQkJDAsMGA0NGDIhHCEyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMv/AABEIAJYAlgMBIgACEQEDEQH/xAGiAAABBQEBAQEBAQAAAAAAAAAAAQIDBAUGBwgJCgsQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+gEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoLEQACAQIEBAMEBwUEBAABAncAAQIDEQQFITEGEkFRB2FxEyIygQgUQpGhscEJIzNS8BVictEKFiQ04SXxFxgZGiYnKCkqNTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqCg4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2dri4+Tl5ufo6ery8/T19vf4+fr/2gAMAwEAAhEDEQA/APn+iiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKMUDg04ldgGDuz19qASG0UUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAf/9k=`

type ImageProcessingServiceTestSuite struct {
	suite.Suite

	service *ImageProcessingService
}

func (suite *ImageProcessingServiceTestSuite) SetupTest() {
	suite.service = NewImageProcessingService()
}

func (suite *ImageProcessingServiceTestSuite) Test_CropCentered_CorrectImagePassed_CroppedCentered() {
	height := 250
	width := 250

	resultHeight := 150
	resultWidth := 150

	testColor := color.RGBA{255, 0, 0, 255}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	img.Set(125, 125, testColor)

	buffer := new(bytes.Buffer)
	jpeg.Encode(buffer, img, nil)

	data, _, _ :=suite.service.CropCentered(buffer.Bytes(), resultWidth, resultHeight, CropTypeDefault)

	encoded := base64.StdEncoding.EncodeToString(data)

	suite.Assert().Equal(encoded, croppedImg)
}

func (suite *ImageProcessingServiceTestSuite) Test_CropCentered_SizePassed_ResultImageWithCorrectSize() {
	height := 250
	width := 250

	resultHeight := 150
	resultWidth := 150
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	buffer := new(bytes.Buffer)
	jpeg.Encode(buffer, img, nil)

	data, _, _ :=suite.service.CropCentered(buffer.Bytes(), resultWidth, resultHeight, CropTypeDefault)

	resizedBuffer := bytes.NewBuffer(data)

	cfg, _ := jpeg.DecodeConfig(resizedBuffer)

	suite.Assert().Equal(cfg.Height, resultHeight)
	suite.Assert().Equal(cfg.Width, resultWidth)
}

func (suite *ImageProcessingServiceTestSuite) Test_CropCentered_SizePassedButBiggerThanImage_ImageSizeNotChanged() {
	height := 100
	width := 100

	resultHeight := 150
	resultWidth := 150
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	buffer := new(bytes.Buffer)
	jpeg.Encode(buffer, img, nil)

	data, _, _ :=suite.service.CropCentered(buffer.Bytes(), resultWidth, resultHeight, CropTypeDefault)

	resizedBuffer := bytes.NewBuffer(data)

	cfg, _ := jpeg.DecodeConfig(resizedBuffer)

	suite.Assert().Equal(cfg.Height, height)
	suite.Assert().Equal(cfg.Width, width)
}

func (suite *ImageProcessingServiceTestSuite) Test_CropCentered_TypeRatio_SizeChangedCorrectly() {
	height := 160
	width := 160

	resultHeight := 16
	resultWidth := 9
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	buffer := new(bytes.Buffer)
	jpeg.Encode(buffer, img, nil)

	data, _, _ :=suite.service.CropCentered(buffer.Bytes(), resultWidth, resultHeight, CropTypeRatio)

	resizedBuffer := bytes.NewBuffer(data)

	cfg, _ := jpeg.DecodeConfig(resizedBuffer)

	suite.Assert().Equal(cfg.Height, 160)
	suite.Assert().Equal(cfg.Width, 90)
}

func (suite *ImageProcessingServiceTestSuite) Test_CropCentered_JPEGPassed_JPEGImageTypeInResult() {
	height := 250
	width := 250

	resultHeight := 150
	resultWidth := 150
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	buffer := new(bytes.Buffer)
	jpeg.Encode(buffer, img, nil)

	_, imageType, _ :=suite.service.CropCentered(buffer.Bytes(), resultWidth, resultHeight, CropTypeDefault)

	suite.Assert().Equal("jpeg", imageType)
}

func (suite *ImageProcessingServiceTestSuite) Test_CropCentered_PNGPassed_PNGImageTypeInResult() {
	height := 250
	width := 250

	resultHeight := 150
	resultWidth := 150
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	buffer := new(bytes.Buffer)
	png.Encode(buffer, img)

	_, imageType, _ :=suite.service.CropCentered(buffer.Bytes(), resultWidth, resultHeight, CropTypeDefault)

	suite.Assert().Equal("png", imageType)
}

func (suite *ImageProcessingServiceTestSuite) Test_CropCentered_GIFPassed_GIFImageTypeInResult() {
	height := 250
	width := 250

	resultHeight := 150
	resultWidth := 150
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	buffer := new(bytes.Buffer)
	gif.Encode(buffer, img, nil)

	_, imageType, _ :=suite.service.CropCentered(buffer.Bytes(), resultWidth, resultHeight, CropTypeDefault)

	suite.Assert().Equal("gif", imageType)
}

func TestImageProcessingService(t *testing.T) {
	testSuite := ImageProcessingServiceTestSuite{}
	suite.Run(t, &testSuite)
}
