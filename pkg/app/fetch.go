package app

import (
	"io/ioutil"
	"net/http"
	"time"
)

type FetchService struct {
	client http.Client
}

func (service *FetchService) GetBodyFromUrl(url string) ([]byte, error) {
	resp, err := service.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func NewFetchService(timeoutMilliseconds int) *FetchService {
	timeout := time.Duration(time.Duration(timeoutMilliseconds) * time.Millisecond)
	client := http.Client{
		Timeout: timeout,
	}

	service := FetchService{}
	service.client = client

	return &service
}
