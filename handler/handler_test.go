package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTreeHandlerOK(t *testing.T) {
	req, err := http.NewRequest("GET", "/?favoriteTree=baobab", nil)
	if err != nil {
		t.Fatal(err)
	}

	tmpl := CreateTemplate()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TemplateHandler(tmpl))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := fmt.Sprintf("It's nice to know that your favorite tree is a <strong>%s</strong>", "baobab")
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestTreeHandlerNOK(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	tmpl := CreateTemplate()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TemplateHandler(tmpl))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Please tell me your favorite tree"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
