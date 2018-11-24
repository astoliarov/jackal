package interfaces

type CropType int

const CropTypeDefault CropType = 1
const CropTypeRatio CropType = 2

//go:generate mockgen -destination=../../tests/mocks/ifetchservice_mock.go -package=mocks jackal/pkg/interfaces IFetchService
type IFetchService interface {
	GetBodyFromUrl(url string) ([]byte, error)
}

//go:generate mockgen -destination=../../tests/mocks/iimageprocessingservice_mock.go -package=mocks jackal/pkg/interfaces IImageProcessingService
type IImageProcessingService interface {
	CropCentered(content []byte, width int, height int, cropType CropType) ([]byte, string, error)
}

//go:generate mockgen -destination=../../tests/mocks/idownloadandcropimageusecase_mock.go -package=mocks jackal/pkg/interfaces IDownloadAndCropImageUseCase
type IDownloadAndCropImageUseCase interface {
	Execute(imageUrl string, width int, height int, cropType CropType) ([]byte, string, error)
}
