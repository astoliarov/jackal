package app

type DownloadAndCropImageUseCase struct {
	fetchService      IFetchService
	processingService IImageProcessingService
}

func (u *DownloadAndCropImageUseCase) Execute(imageUrl string, width int, height int, cropType CropType) ([]byte, string, error) {

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

func NewDownloadAndCropImageUseCase(imageProcessingService IImageProcessingService, fetchService IFetchService) *DownloadAndCropImageUseCase {
	return &DownloadAndCropImageUseCase{fetchService: fetchService, processingService: imageProcessingService}
}
