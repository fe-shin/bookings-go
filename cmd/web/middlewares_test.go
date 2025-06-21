package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var httpHandlerMock httpHandlerMockType
	handler := noSurf(&httpHandlerMock)

	switch responseType := handler.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("type is not http Handler, but instead it's %T", responseType)
	}
}

func TestSessionLoad(t *testing.T) {
	var httpHandlerMock httpHandlerMockType
	handler := sessionLoad(&httpHandlerMock)

	switch responseType := handler.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("type is not http Handler, but instead it's %T", responseType)
	}
}
