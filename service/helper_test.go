package service

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/helphone/api/manager"
)

func TestRenderWithETag(t *testing.T) {
	req, w := prepareRequest("test")
	response, _ := manager.GetCountries("en")
	data, _ := json.Marshal(response)

	renderWithETag(w, req, data)

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json; charset=utf-8" {
		t.Errorf("Content-Type is not good. Receive %s", contentType)
	}

	eTag := w.Header().Get("ETag")
	if eTag == "" {
		t.Error("ETag should be set")
	}

	req, w = prepareRequest("test")
	req.Header.Set("If-None-Match", eTag)

	renderWithETag(w, req, data)

	if w.Code != http.StatusNotModified {
		t.Errorf("The status code of the second request with ETag should be not modified. Got %v", w.Code)
	}

	if len(w.Body.String()) > 0 {
		t.Errorf("The body of a not modified status should be empty. Got %s", w.Body.String())
	}
}

func TestPrintANotFoundCountry(t *testing.T) {
	req, w := prepareRequest("test")
	printANotFoundCountry(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("The status code should be not found. Got %v", w.Code)
	}

	responseError := JSONError{}
	err := json.Unmarshal(w.Body.Bytes(), &responseError)
	if err != nil {
		t.Errorf("The Unmarshal goes wrong to invert the encode. Error %v", err)
	}

	if responseError.Message != countryNotFoundMessage {
		t.Errorf("Bad message in the JSON. Receive %s", responseError.Message)
	}
}
