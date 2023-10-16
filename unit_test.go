package main

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKeyValid(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-api-key")
	expected := "my-api-key"
	actual, err := auth.GetAPIKey(headers)
	if err != nil {
		t.Errorf("TestGetAPIKeyValid failed: expected no error but got %v", err)
	}
	if actual != expected {
		t.Errorf("TestGetAPIKeyValid failed: expected %s but got %s", expected, actual)
	}
}

func TestGetAPIKeyInvalid(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer my-token")
	expected := "b"
	actual, err := auth.GetAPIKey(headers)
	if err == nil {
		t.Errorf("TestGetAPIKeyInvalid failed: expected error but got no error")
	}
	if actual != expected {
		t.Errorf("TestGetAPIKeyInvalid failed: expected %s but got %s", expected, actual)
	}
}
