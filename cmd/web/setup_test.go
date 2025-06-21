package main

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

type httpHandlerMockType struct{}

func (handler *httpHandlerMockType) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
