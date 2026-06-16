package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandlerPlaceholder(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Could not create HTTP request: %v", err)
	}

	rec := httptest.NewRecorder()

	_ = req
	_ = rec
}

func TestArtistHandlerNotFoundPlaceholder(t *testing.T) {
	req, err := http.NewRequest("GET", "/artist?id=9999", nil)
	if err != nil {
		t.Fatalf("Could not create HTTP request: %v", err)
	}

	rec := httptest.NewRecorder()

	_ = req
	_ = rec
}