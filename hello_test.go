package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/user"
)

func TestLoggedIn(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}

	w := httptest.NewRecorder()

	aetest.Login(&user.User{Email: "test@example.org"}, req)

	handle(w, req)

	code := w.Code
	if code != http.StatusOK {
		t.Errorf("got code %v, want %v", code, http.StatusOK)
	}

	body := w.Body.Bytes()
	expected := []byte("Hello, world!")
	if !bytes.Contains(body, expected) {
		t.Errorf("got body %v, did not contain %v", string(body), string(expected))
	}
}

func TestNotLoggedIn(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}

	w := httptest.NewRecorder()

	handle(w, req)

	code := w.Code
	if code != http.StatusForbidden {
		t.Errorf("got code %v, want %v", code, http.StatusForbidden)
	}

	body := w.Body.Bytes()
	expected := []byte("sorry, not logged in")
	if !bytes.Contains(body, expected) {
		t.Errorf("got body %v, did not contain %v", string(body), string(expected))
	}
}
