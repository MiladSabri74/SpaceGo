package main

import (
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestSample_GetTrafficStatus(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Fatalf("خطا در راه‌اندازی سرور تستی: %v", err)
	}

	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "light"})
	}))
	server.Listener.Close()
	server.Listener = listener
	server.Start()
	t.Cleanup(func() {
		server.Close()
	})

	os.Setenv("API_KEY", "sample_key")
	status, err := GetTrafficStatus("Tehran")

	if err != nil {
		t.Fatalf("در حالت عادی نباید خطایی برگردد، اما خطای زیر رخ داد: %v", err)
	}
	if status != "light" {
		t.Errorf("وضعیت دریافتی اشتباه است. مقدار مورد انتظار: light، مقدار دریافتی: %q", status)
	}
}
