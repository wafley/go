package service

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetStudents(t *testing.T) {
	InitLogger()

	req, err := http.NewRequest(http.MethodGet, "/students", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetStudents)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected application/json, got %s", contentType)
	}

	if rr.Body.Len() == 0 {
		t.Errorf("expected body not empty")
	}
}