package api

import (
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/astoliarov/jackal/pkg/interfaces"
	"github.com/astoliarov/jackal/tests/mocks"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

type CropHandlerTestSuite struct {
	suite.Suite

	router http.Handler

	controller *gomock.Controller

	useCaseMock *mocks.MockIDownloadAndCropImageUseCase
}

func (suite *CropHandlerTestSuite) SetupTest() {
	suite.controller = gomock.NewController(suite.T())
	suite.useCaseMock = mocks.NewMockIDownloadAndCropImageUseCase(suite.controller)

	apiService := NewAPIService(suite.useCaseMock, true)

	suite.router = apiService.GetRouter()
}

func (suite *CropHandlerTestSuite) Test_Handle_AllParametersPassed_CorrectStatus() {
	width := 400
	height := 400
	url := "test"

	path := fmt.Sprintf("/api/v1/crop?width=%d&height=%d&url=%s", width, height, url)
	suite.useCaseMock.EXPECT().Execute(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]byte{}, "jpeg", nil)

	rr := performRequest(suite.router, "GET", path)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)
}

func (suite *CropHandlerTestSuite) Test_Handle_PassedWidth_WidthParsedCorrectly() {
	width := 400
	height := 400
	url := "test"

	path := fmt.Sprintf("/api/v1/crop?width=%d&height=%d&url=%s", width, height, url)
	suite.useCaseMock.EXPECT().Execute(gomock.Any(), width, gomock.Any(), gomock.Any()).Return([]byte{}, "jpeg", nil)

	_ = performRequest(suite.router, "GET", path)

	suite.controller.Finish()
}

func (suite *CropHandlerTestSuite) Test_Handle_PassedHeight_HeightParsedCorrectly() {
	width := 400
	height := 400
	url := "test"

	path := fmt.Sprintf("/api/v1/crop?width=%d&height=%d&url=%s", width, height, url)
	suite.useCaseMock.EXPECT().Execute(gomock.Any(), gomock.Any(), height, gomock.Any()).Return([]byte{}, "jpeg", nil)

	_ = performRequest(suite.router, "GET", path)

	suite.controller.Finish()
}

func (suite *CropHandlerTestSuite) Test_Handle_PassedUrl_UrlParsedCorrectly() {
	width := 400
	height := 400
	urlString := "http://foobarbuz.tech/ololo?bing=test"
	encodedUrl := url.QueryEscape(urlString)

	path := fmt.Sprintf("/api/v1/crop?width=%d&height=%d&url=%s", width, height, encodedUrl)
	suite.useCaseMock.EXPECT().Execute(urlString, gomock.Any(), height, gomock.Any()).Return([]byte{}, "jpeg", nil)

	_ = performRequest(suite.router, "GET", path)

	suite.controller.Finish()
}

func (suite *CropHandlerTestSuite) Test_Handle_CropTypeNotPassed_CropTypeDefaultSelected() {
	width := 400
	height := 400
	urlString := "http://foobarbuz.tech/ololo?bing=test"
	encodedUrl := url.QueryEscape(urlString)

	path := fmt.Sprintf("/api/v1/crop?width=%d&height=%d&url=%s", width, height, encodedUrl)
	suite.useCaseMock.EXPECT().Execute(urlString, gomock.Any(), gomock.Any(), interfaces.CropTypeDefault).Return([]byte{}, "jpeg", nil)

	_ = performRequest(suite.router, "GET", path)

	suite.controller.Finish()
}

func (suite *CropHandlerTestSuite) Test_Handle_CropTypePassedRation_CropTypeRatioSelected() {
	width := 400
	height := 400
	urlString := "http://foobarbuz.tech/ololo?bing=test"
	encodedUrl := url.QueryEscape(urlString)

	path := fmt.Sprintf("/api/v1/crop?type=ratio&width=%d&height=%d&url=%s", width, height, encodedUrl)
	suite.useCaseMock.EXPECT().Execute(urlString, gomock.Any(), gomock.Any(), interfaces.CropTypeRatio).Return([]byte{}, "jpeg", nil)

	_ = performRequest(suite.router, "GET", path)

	suite.controller.Finish()
}

func (suite *CropHandlerTestSuite) Test_Handle_CropTypePassedNotRatio_CropTypeDefaultSelected() {
	width := 400
	height := 400
	urlString := "http://foobarbuz.tech/ololo?bing=test"
	encodedUrl := url.QueryEscape(urlString)

	path := fmt.Sprintf("/api/v1/crop?type=pration&width=%d&height=%d&url=%s", width, height, encodedUrl)
	suite.useCaseMock.EXPECT().Execute(urlString, gomock.Any(), gomock.Any(), interfaces.CropTypeDefault).Return([]byte{}, "jpeg", nil)

	_ = performRequest(suite.router, "GET", path)

	suite.controller.Finish()
}

func (suite *CropHandlerTestSuite) Test_Handle_ParamsNotPassed_StatusCorrect() {
	rr := performRequest(suite.router, "GET", "/api/v1/crop")

	assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
}

func (suite *CropHandlerTestSuite) Test_Handle_WidthNotPassed_WidthErrorMessageCorrect() {
	rr := performRequest(suite.router, "GET", "/api/v1/crop?height=400&url=test")

	var responseJson map[string]map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &responseJson)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), responseJson["errors"]["width"], "field is required")
}

func (suite *CropHandlerTestSuite) Test_Handle_WidthPassedButNotCorrect_WidthErrorMessageCorrect() {
	rr := performRequest(suite.router, "GET", "/api/v1/crop?height=400&url=test&width=test")

	var responseJson map[string]map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &responseJson)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), responseJson["errors"]["width"], "bad value")
}

func (suite *CropHandlerTestSuite) Test_Handle_HeightNotPassed_HeightErrorMessageCorrect() {
	rr := performRequest(suite.router, "GET", "/api/v1/crop?width=400&url=test")

	var responseJson map[string]map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &responseJson)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), responseJson["errors"]["height"], "field is required")
}

func (suite *CropHandlerTestSuite) Test_Handle_HeightPassedButNotCorrect_HeightErrorMessageCorrect() {
	rr := performRequest(suite.router, "GET", "/api/v1/crop?height=test&url=test&width=400")

	var responseJson map[string]map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &responseJson)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), responseJson["errors"]["height"], "bad value")
}

func (suite *CropHandlerTestSuite) Test_Handle_UrlNotPassed_UrlErrorMessageCorrect() {
	rr := performRequest(suite.router, "GET", "/api/v1/crop?width=400&height=400")

	var responseJson map[string]map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &responseJson)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), responseJson["errors"]["url"], "field is required")
}

func TestCropHandler(t *testing.T) {
	testSuite := CropHandlerTestSuite{}
	suite.Run(t, &testSuite)
}
