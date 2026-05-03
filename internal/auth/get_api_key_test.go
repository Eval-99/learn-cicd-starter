package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeyValid(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://blargus.com", nil)

	token := "Butt"
	req.Header.Set("Authorization", "ApiKey "+token)

	key, err := GetAPIKey(req.Header)
	if err != nil {
		t.Fatalf("Error extracting API key from header: %v", err)
	}

	if key != token {
		t.Fatalf("expected: %v, got: %v", token, key)
	}
}

func TestGetAPIKeyInvalid(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://blargus.com", nil)

	_, err := GetAPIKey(req.Header)
	if err == nil {
		t.Fatalf("Error GetAPIKey did't fail with faulty header: %v", err)
	}
}
