package analyser

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockCrawler struct {
	response map[string]interface{}
	err      error
}

func (m *MockCrawler) FromURL(url string) (map[string]interface{}, error) {
	return m.response, m.err
}

func TestAnalyserPostPositive(t *testing.T) {
	responseMap := make(map[string]interface{})
	responseMap["data"] = "something"
	mockCrawler := &MockCrawler{}
	mockCrawler.response = responseMap

	analyser := Analyser{
		Crawler: mockCrawler,
	}

	request := httptest.NewRequest("POST", "/api/v1/analyser/html", strings.NewReader(`
	{"data": "https://www.google.co.in"}
	`))

	recorder := httptest.NewRecorder()
	analyser.Post(recorder, request)
	result := recorder.Result()
	if result.StatusCode != 200 {
		t.Errorf("Positive test case if failing")
	}
}

func TestAnalyserPostNegative(t *testing.T) {
	mockCrawler := &MockCrawler{}
	mockCrawler.err = errors.New("no such host")
	analyser := Analyser{
		Crawler: mockCrawler,
	}

	requestMethodNotAllowed := httptest.NewRequest("GET", "/api/v1/analyser/html", nil)
	recorder := httptest.NewRecorder()
	analyser.Post(recorder, requestMethodNotAllowed)
	result := recorder.Result()
	if result.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("StatusMethodNotAllowed test case failing with %v Expected: %v", result.StatusCode, http.StatusMethodNotAllowed)
	}

	requestUnsupportedMediaType := httptest.NewRequest("POST", "/api/v1/analyser/html", strings.NewReader(`
	{"data": "https://www.google.co.in"
	`))
	recorder = httptest.NewRecorder()
	analyser.Post(recorder, requestUnsupportedMediaType)
	result = recorder.Result()
	if result.StatusCode != http.StatusUnsupportedMediaType {
		t.Errorf("StatusUnsupportedMediaType test case failing with %v Expected: %v", result.StatusCode, http.StatusUnsupportedMediaType)
	}

	requestStatusUnprocessableEntity := httptest.NewRequest("POST", "/api/v1/analyser/html", strings.NewReader(`
	{"data": "https://www.google.co.in"}
	`))
	recorder = httptest.NewRecorder()
	analyser.Post(recorder, requestStatusUnprocessableEntity)
	result = recorder.Result()
	if result.StatusCode != http.StatusUnprocessableEntity {
		t.Errorf("StatusUnprocessableEntity test case failing with %v Expected: %v", result.StatusCode, http.StatusUnprocessableEntity)
	}

	mockCrawler.err = errors.New("hello world parse error")
	requestStatusInternalServerError := httptest.NewRequest("POST", "/api/v1/analyser/html", strings.NewReader(`
	{"data": "https://www.google.co.in"}
	`))
	recorder = httptest.NewRecorder()
	analyser.Post(recorder, requestStatusInternalServerError)
	result = recorder.Result()
	if result.StatusCode != http.StatusInternalServerError {
		t.Errorf("UnsupportedMediaType test case failing with %v Expected: %v", result.StatusCode, http.StatusInternalServerError)
	}

}
