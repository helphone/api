package service

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/helphone/api/manager"
	"github.com/helphone/api/model"
)

func init() {
	manager.Init()
	<-manager.ReadyChannel
}

func prepareRequest(path string) (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest("GET", "http://example.com"+path, nil)
	w := httptest.NewRecorder()
	return req, w
}

func TestPing(t *testing.T) {
	req, w := prepareRequest("/ping")
	ping(w, req)

	if w.Code != 200 {
		t.Error("Result of ping should return a 200")
	}

	body := w.Body.String()
	if body != "Pong" {
		t.Errorf("Body of ping should return be 'Pong'. Receive '%s'", body)
	}
}

func TestGetCountries(t *testing.T) {
	req, w := prepareRequest("/countries")
	getCountries(w, req)

	if w.Code != 200 {
		t.Error("Result of ping should return a 200")
	}

	body := w.Body.String()
	if len(body) < 2 {
		t.Errorf("Body length should be at least a empty JSON. Receive length %v", len(body))
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json; charset=utf-8" {
		t.Errorf("Content-Type should be a JSON. Receive '%s'", contentType)
	}

	response := model.Response{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("The Unmarshal goes wrong to invert the encode. Error %s", err)
	}

	req, w = prepareRequest("/countries?lat=100.0&long=100.0")
	getPhonenumbers(w, req)

	if w.Code != http.StatusNotFound {
		t.Error("Result of the request should return a not found status")
	}

	body = w.Body.String()
	if len(body) < 2 {
		t.Errorf("Body length should be at least a empty JSON. Receive length %v", len(body))
	}

	req, w = prepareRequest("/countries?lat=50.625073&long=4.746094")
	getPhonenumbers(w, req)

	if w.Code != http.StatusOK {
		t.Error("Result of the request should return an ok status")
	}

	contentType = w.Header().Get("Content-Type")
	if contentType != "application/json; charset=utf-8" {
		t.Errorf("Content-Type should be a JSON. Receive '%s'", contentType)
	}

	body = w.Body.String()
	if len(body) < 2 {
		t.Errorf("Body length should be at least a empty JSON. Receive length %v", len(body))
	}
}

func TestGetPhonenumbers(t *testing.T) {
	req, w := prepareRequest("/phonenumbers?language=fr")
	getPhonenumbers(w, req)

	if w.Code != 200 {
		t.Error("Result of the request should return a 200")
	}

	body := w.Body.String()
	if len(body) < 2 {
		t.Errorf("Body length should be at least a empty JSON. Receive length %v", len(body))
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json; charset=utf-8" {
		t.Errorf("Content-Type should be a JSON. Receive '%s'", contentType)
	}

	response := model.Response{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("The Unmarshal goes wrong to invert the encode. Error %s", err)
	}

	req, w = prepareRequest("/phonenumbers?country=be")
	getPhonenumbers(w, req)

	if w.Code != 200 {
		t.Error("Result of the request should return a 200")
	}

	body = w.Body.String()
	if len(body) < 2 {
		t.Errorf("Body length should be at least a empty JSON. Receive length %v", len(body))
	}

	response = model.Response{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("The Unmarshal goes wrong to invert the encode. Error %v", err)
	}

	req, w = prepareRequest("/phonenumbers?lat=10.0&language=fr")
	getPhonenumbers(w, req)

	if w.Code != http.StatusBadRequest {
		t.Error("Result of the request should return a bad request status")
	}

	body = w.Body.String()
	if len(body) < 2 {
		t.Errorf("Body length should be at least a empty JSON. Receive length %v", len(body))
	}

	responseError := JSONError{}
	err = json.Unmarshal(w.Body.Bytes(), &responseError)
	if err != nil {
		t.Errorf("The Unmarshal goes wrong to invert the encode. Error %v", err)
	}

	req, w = prepareRequest("/phonenumbers?lat=100.0&long=100.0")
	getPhonenumbers(w, req)

	if w.Code != http.StatusNotFound {
		t.Error("Result of the request should return a not found status")
	}

	body = w.Body.String()
	if len(body) < 2 {
		t.Errorf("Body length should be at least a empty JSON. Receive length %v", len(body))
	}

	responseError = JSONError{}
	err = json.Unmarshal(w.Body.Bytes(), &responseError)
	if err != nil {
		t.Errorf("The Unmarshal goes wrong to invert the encode. Error %v", err)
	}
}
