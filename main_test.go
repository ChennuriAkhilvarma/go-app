package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()

	HelloHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	body := w.Body.String()
	if !strings.Contains(body, "Hello, DevOps World!") {
		t.Errorf("Expected body to contain greeting, got %s", body)
	}
}

func TestHelloHandler_StatusOK(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	rec := httptest.NewRecorder()

	HelloHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestHelloHandler_BodyExactMatch(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	rec := httptest.NewRecorder()

	HelloHandler(rec, req)

	expected := "Hello, DevOps World!"
	actual := rec.Body.String()

	if actual != expected {
		t.Errorf("Expected response body %q, got %q", expected, actual)
	}
}
