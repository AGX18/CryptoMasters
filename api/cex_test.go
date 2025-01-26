package api_test

import (
	"testing"

	"github.com/AGX18/CryptoMasters/api"
)

func TestAPICall(t *testing.T) {
	_, err := api.GetRate("")

	if err == nil {
		t.Error("error was not found")
	}
	
}