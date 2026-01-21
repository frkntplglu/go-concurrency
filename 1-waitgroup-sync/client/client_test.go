package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type MockHTTPClient struct {
	GetFunc func(ctx context.Context, url string, headers map[string]string) ([]byte, error)
}

func (m *MockHTTPClient) Get(ctx context.Context, url string, headers map[string]string) ([]byte, error) {
	return m.GetFunc(ctx, url, headers)
}

func TestClient_Get_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		if auth := r.Header.Get("Authorization"); auth != "Bearer test-token" {
			t.Errorf("Expected Authorization header, got %s", auth)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"id":   1,
			"name": "Test User",
		})
	}))
	defer server.Close()

	client := New(5 * time.Second)

	headers := map[string]string{
		"Authorization": "Bearer test-token",
	}
	_, err := client.Get(context.Background(), server.URL, headers)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}

}

func TestClient_Get_Timeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second) // 3 saniye bekle
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := New(1 * time.Second)

	_, err := client.Get(context.Background(), server.URL, nil)
	if err == nil {
		t.Error("Expected timeout error, got nil")
	}
}

func TestClient_Get_ContextCancel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := New(5 * time.Second)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := client.Get(ctx, server.URL, nil)
	if err == nil {
		t.Error("Expected context canceled error, got nil")
	}
}

func TestClient_Get_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "not found"}`))
	}))
	defer server.Close()

	client := New(5 * time.Second)

	_, err := client.Get(context.Background(), server.URL, nil)

	if err != nil {
		t.Logf("Got error as expected: %v", err)
	}

}
