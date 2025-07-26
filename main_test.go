package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test using httptest
func TestTasksIntegration(t *testing.T) {
	server := httptest.NewServer(setupRouter())
	defer server.Close()

	resp, err := http.Get(server.URL + "/tasks")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Got non ok status code: %d", resp.StatusCode)
	}
}
