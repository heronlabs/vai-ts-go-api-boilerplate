package healthCheckController

import (
	"bytes"
	"encoding/json"
	healthCheckController "hello-world/src/application/api/controllers/health-check"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetIndex(t *testing.T) {
	var expectedMessage = "Server running!"

	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(healthCheckController.GetIndex)
	handler.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]string
	if err := json.NewDecoder(responseRecorder.Body).Decode(&response); err != nil {
		t.Errorf("Error decoding JSON response: %v", err)
	}

	if payload, ok := response["payload"]; !ok || payload != expectedMessage {
		t.Errorf("Unexpected field in JSON response: got %v want '%v'", payload, expectedMessage)
	}
}

func TestAnyIndex(t *testing.T) {
	request, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(healthCheckController.GetIndex)
	handler.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("Expected MethodNotAllowed")
	}
}

func TestPostWebHook(t *testing.T) {
	body := struct {
		TestOne string `json:"testOne"`
	}{
		TestOne: "hello world!",
	}

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/webhook", bytes.NewReader(jsonBytes))
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(healthCheckController.PostWebHook)

	handler.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestPostWebHookWrongBody(t *testing.T) {
	request, err := http.NewRequest("POST", "/webhook", bytes.NewReader(nil))
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(healthCheckController.PostWebHook)

	handler.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		t.Errorf("Expected BadRequest")
	}
}

func TestAnyWebHook(t *testing.T) {

	request, err := http.NewRequest("GET", "/webhook", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(healthCheckController.PostWebHook)

	handler.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("Expected MethodNotAllowed")
	}
}
