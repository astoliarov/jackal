package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jackal/pkg/app"
	"log"
	"net/http"
)

type ValidateErrs map[string]string

func (v *ValidateErrs) IsEmpty() bool {
	return len(*v) == 0
}

type APIService struct {
	router *gin.Engine
}

func (s *APIService) Run(port int) {
	err := s.router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
}

func (s *APIService) GetRouter() http.Handler {
	return s.router
}

func NewAPIService(cropUseCase app.IDownloadAndCropImageUseCase, release bool) *APIService {
	apiService := APIService{}
	apiService.router = gin.Default()

	if release {
		gin.SetMode(gin.ReleaseMode)
	}

	cropHandler := NewCropHandler(cropUseCase)

	apiService.router.GET("/api/v1/crop", cropHandler.HandleGetCrop)

	return &apiService
}
