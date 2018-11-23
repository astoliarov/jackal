package main

import (
	"jackal/pkg/api"
	"jackal/pkg/app"
	"log"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	apiService := api.NewAPIService(app.GetDownloadAndCropService(), true)
	apiService.Run(app.GetConfig().Port)
}
