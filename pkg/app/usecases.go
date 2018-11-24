package app

import "jackal/pkg/interfaces"

type DownloadAndCropImageUseCase struct {
	fetchService      interfaces.IFetchService
	processingService interfaces.IImageProcessingService
}

func (u *DownloadAndCropImageUseCase) Execute(imageUrl string, width int, height int, cropType interfaces.CropType) ([]byte, string, error) {

	content, err := u.fetchService.GetBodyFromUrl(imageUrl)
	if err != nil {
		return nil, "", err
	}

	croppedContent, imageType, err := u.processingService.CropCentered(content, width, height, cropType)
	if err != nil {
		return nil, "", err
	}

	return croppedContent, imageType, nil
}

func NewDownloadAndCropImageUseCase(imageProcessingService interfaces.IImageProcessingService, fetchService interfaces.IFetchService) *DownloadAndCropImageUseCase {
	return &DownloadAndCropImageUseCase{fetchService: fetchService, processingService: imageProcessingService}
}
