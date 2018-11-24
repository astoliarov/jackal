package app

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"jackal/pkg/interfaces"
	"jackal/tests/mocks"
	"testing"
)

type DownloadAndCropImageUseCaseTestSuite struct {
	suite.Suite

	controller            *gomock.Controller
	fetchServiceMock      *mocks.MockIFetchService
	processingServiceMock *mocks.MockIImageProcessingService

	useCase *DownloadAndCropImageUseCase
}

func (suite *DownloadAndCropImageUseCaseTestSuite) SetupTest() {
	suite.controller = gomock.NewController(suite.T())

	suite.fetchServiceMock = mocks.NewMockIFetchService(suite.controller)
	suite.processingServiceMock = mocks.NewMockIImageProcessingService(suite.controller)

	suite.useCase = NewDownloadAndCropImageUseCase(suite.processingServiceMock, suite.fetchServiceMock)
}

func (suite *DownloadAndCropImageUseCaseTestSuite) Test_Execute_GetBodyFromUrlReturnErr_ReturnErr() {

	dummyErr := errors.New("dummy err")

	suite.fetchServiceMock.EXPECT().GetBodyFromUrl(gomock.Any()).Return([]byte{}, dummyErr)

	result, imgType, err := suite.useCase.Execute("", 0, 0, interfaces.CropTypeDefault)

	suite.Assert().Equal("", imgType)
	suite.Assert().Nil(result)
	suite.Assert().Equal(dummyErr, err)
}

func (suite *DownloadAndCropImageUseCaseTestSuite) Test_Execute_UrlPassed_UrlPassedCorrectlyToGetBodyFromUrl() {

	dummyErr := errors.New("dummy err")
	url := "http://test.com"

	suite.fetchServiceMock.EXPECT().GetBodyFromUrl(url).Return([]byte{}, dummyErr)

	suite.useCase.Execute(url, 0, 0, interfaces.CropTypeDefault)

	suite.controller.Finish()
}

func (suite *DownloadAndCropImageUseCaseTestSuite) Test_Execute_ImgDownloaded_ImgBytesPassedToCropCentered() {

	imgData := []byte("test")

	suite.fetchServiceMock.EXPECT().GetBodyFromUrl(gomock.Any()).Return(imgData, nil)
	suite.processingServiceMock.EXPECT().CropCentered(imgData, gomock.Any(), gomock.Any(), gomock.Any()).Return(imgData, "", nil)

	suite.useCase.Execute("", 0, 0, interfaces.CropTypeDefault)

	suite.controller.Finish()
}

func (suite *DownloadAndCropImageUseCaseTestSuite) Test_Execute_HeightPassed_HeightPassedCorrectlyToCropCentered() {

	imgData := []byte("test")
	height := 13

	suite.fetchServiceMock.EXPECT().GetBodyFromUrl(gomock.Any()).Return(imgData, nil)
	suite.processingServiceMock.EXPECT().CropCentered(gomock.Any(), gomock.Any(), height, gomock.Any()).Return(imgData, "", nil)

	suite.useCase.Execute("", 0, height, interfaces.CropTypeDefault)

	suite.controller.Finish()
}

func (suite *DownloadAndCropImageUseCaseTestSuite) Test_Execute_WidthPassed_WidthPassedCorrectlyToCropCentered() {

	imgData := []byte("test")
	width := 13

	suite.fetchServiceMock.EXPECT().GetBodyFromUrl(gomock.Any()).Return(imgData, nil)
	suite.processingServiceMock.EXPECT().CropCentered(gomock.Any(), width, gomock.Any(), gomock.Any()).Return(imgData, "", nil)

	suite.useCase.Execute("", width, 0, interfaces.CropTypeDefault)

	suite.controller.Finish()
}

func (suite *DownloadAndCropImageUseCaseTestSuite) Test_Execute_CropTypePassed_CropTypePassedCorrectlyToCropCentered() {

	imgData := []byte("test")

	suite.fetchServiceMock.EXPECT().GetBodyFromUrl(gomock.Any()).Return(imgData, nil)
	suite.processingServiceMock.EXPECT().CropCentered(gomock.Any(), gomock.Any(), gomock.Any(), interfaces.CropTypeRatio).Return(imgData, "", nil)

	suite.useCase.Execute("", 0, 0, interfaces.CropTypeRatio)

	suite.controller.Finish()
}

func (suite *DownloadAndCropImageUseCaseTestSuite) Test_Execute_ParamsPassedCorrectly_ReturnedDataCorrect() {

	imgData := []byte("test")
	imgType := "jpeg"

	suite.fetchServiceMock.EXPECT().GetBodyFromUrl(gomock.Any()).Return(imgData, nil)
	suite.processingServiceMock.EXPECT().CropCentered(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(imgData, imgType, nil)

	data, resultType, err := suite.useCase.Execute("", 0, 0, interfaces.CropTypeRatio)

	suite.Assert().Equal(imgData, data)
	suite.Assert().Equal(imgType, resultType)
	suite.Assert().Nil(err)

}

func (suite *DownloadAndCropImageUseCaseTestSuite) Test_Execute_CropCenteredReturnError_ReturnedDataCorrect() {

	dummyErr := errors.New("dummy err")

	suite.fetchServiceMock.EXPECT().GetBodyFromUrl(gomock.Any()).Return([]byte{}, nil)
	suite.processingServiceMock.EXPECT().CropCentered(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]byte{}, "test", dummyErr)

	data, resultType, err := suite.useCase.Execute("", 0, 0, interfaces.CropTypeRatio)

	suite.Assert().Equal([]byte(nil), data)
	suite.Assert().Equal("", resultType)
	suite.Assert().Equal(dummyErr, err)

}

func TestDownloadAndCropImageUseCase(t *testing.T) {
	testSuite := DownloadAndCropImageUseCaseTestSuite{}
	suite.Run(t, &testSuite)
}
