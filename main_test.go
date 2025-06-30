package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Real server, not httptest
func startTestServer() *http.Server {
	mux := setupRouter()
	srv := &http.Server{Addr: ":0", Handler: mux} // :0 tells OS to assign a free port

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	return srv
}

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

func TestRealServerTasksIntegration(t *testing.T) {
	server := startTestServer()

	resp, err := http.Get("http://" + server.Addr + "/tasks")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Got non ok status code: %d", resp.StatusCode)
	}
}
