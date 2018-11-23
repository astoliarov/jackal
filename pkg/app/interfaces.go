package app

type IFetchService interface {
	GetBodyFromUrl(url string) ([]byte, error)
}

type IImageProcessingService interface {
	CropCentered(content []byte, width int, height int, cropType CropType) ([]byte, string, error)
}

//go:generate mockgen -destination=../../tests/mocks/idownloadandcropimageusecase_mock.go -package=mocks jackal/pkg/app IDownloadAndCropImageUseCase
type IDownloadAndCropImageUseCase interface {
	Execute(imageUrl string, width int, height int, cropType CropType) ([]byte, string, error)
}
