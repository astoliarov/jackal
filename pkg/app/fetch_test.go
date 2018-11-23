package app

import (
	"testing"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net"
	"fmt"
	"context"
	"time"
	"net/url"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func timeOutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))

	time.Sleep(2 * time.Second)
}


type FetchServiceTestSuite struct {
	suite.Suite

	server *http.Server
	listener net.Listener
	fetchService *FetchService
}

func (suite *FetchServiceTestSuite) SetupTest() {
	suite.fetchService = NewFetchService(1000)

	mux := http.NewServeMux()
	mux.HandleFunc("/hw", helloWorldHandler)
	mux.HandleFunc("/to", timeOutHandler)

	suite.server = &http.Server{Handler: mux}
	suite.listener = suite.startListenOnRandomPort()
}


func (suite *FetchServiceTestSuite) startListenOnRandomPort() net.Listener {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	return listener
}

func (suite *FetchServiceTestSuite) runServer(ctx context.Context){
	go func() {
		<- ctx.Done()
		suite.server.Shutdown(ctx)
	}()

	go func() {
		suite.server.Serve(suite.listener)
	}()
}

func (suite *FetchServiceTestSuite) Test_GetBodyFromUrl_GotCorrectUrl_ReturnCorrectResult() {

	ctx, shutdownServer := context.WithCancel(context.Background())
	suite.runServer(ctx)

	imgUrl := fmt.Sprintf("http://%s%s", suite.listener.Addr().String(), "/hw")

	data, err := suite.fetchService.GetBodyFromUrl(imgUrl)

	shutdownServer()

	suite.Assert().Nil(err)
	suite.Assert().Equal(string(data), "hello world")
}

func (suite *FetchServiceTestSuite) Test_GetBodyFromUrl_TimedOut_ReturnErr() {

	ctx, shutdownServer := context.WithCancel(context.Background())
	suite.runServer(ctx)
	imgUrl := fmt.Sprintf("http://%s%s", suite.listener.Addr().String(), "/to")

	_, err := suite.fetchService.GetBodyFromUrl(imgUrl)

	shutdownServer()

	urlErr := err.(*url.Error)
	netErr := urlErr.Err.(net.Error)

	suite.Assert().True(netErr.Timeout())
}

func TestFetchService(t *testing.T) {
	testSuite := FetchServiceTestSuite{}
	suite.Run(t, &testSuite)
}

