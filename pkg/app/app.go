package app

type App struct {
	fetchService           IFetchService
	imageProcessingService IImageProcessingService

	downloadAndCropUseCase *DownloadAndCropImageUseCase

	config *Config
}

func (app *App) GetDownloadAndCropService() *DownloadAndCropImageUseCase {
	return app.downloadAndCropUseCase
}

func (app *App) GetConfig() *Config {
	return app.config
}

func NewApp() (*App, error) {
	app := App{}

	config, err := ConfitaConfigLoader()
	if err != nil {
		return nil, err
	}

	app.config = config
	app.fetchService = NewFetchService(app.config.FetchTimeout)
	app.imageProcessingService = NewImageProcessingService()

	app.downloadAndCropUseCase = NewDownloadAndCropImageUseCase(app.imageProcessingService, app.fetchService)

	return &app, nil
}
