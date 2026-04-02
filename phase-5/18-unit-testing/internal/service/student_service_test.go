package service

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetStudents(t *testing.T) {
	InitLogger()

	originalSleep := sleep
	defer func() { sleep = originalSleep }()

	sleep = func(d time.Duration) {}

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

func TestCalculateGrade(t *testing.T) {
    tests := []struct {
        name     string
        score    int
        expected string
    }{
        {"Score A", 95, "A"},
        {"Score B", 85, "B"},
        {"Score C", 75, "C"},
        {"Score D", 60, "D"},
        {"Limit 90", 90, "A"},
        {"Limit 80", 80, "B"},
		{"Very Low Score", 0, "D"},
		{"Negative Score", -10, "D"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := CalculateGrade(tt.score)
            if result != tt.expected {
                t.Errorf("%s FAILED: score %d expected %s, but got %s", tt.name, tt.score, tt.expected, result)
            }
        })
    }
}