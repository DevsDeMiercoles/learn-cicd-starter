package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	var emptyheader http.Header = make(map[string][]string)

	_, err := GetAPIKey(emptyheader)
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("Empty header, expected error: %v, got %v", ErrNoAuthHeaderIncluded, err)
	}

	malformedHeader := emptyheader
	malformedHeader.Add("Authorization", "ApiToken")
	_, err = GetAPIKey(emptyheader)
	if err == nil {
		t.Errorf("Malformed header, expected error")
	}

	malformedHeader.Set("Authorization", "ApiKey")
	_, err = GetAPIKey(emptyheader)
	if err == nil {
		t.Errorf("Malformed header, expected error")
	}

	properHeader := malformedHeader
	expectedApiKey := "The-Key"
	properHeader.Set("Authorization", "ApiKey "+expectedApiKey)
	apiKey, err := GetAPIKey(emptyheader)
	if err != nil {
		t.Errorf("Proper Header, unexpected error: %v", err)
	}
	if apiKey != expectedApiKey {
		t.Errorf("Proper header, the apikey was extracted wrongly. Expected %v | Got %v", expectedApiKey, apiKey)
	}
}
