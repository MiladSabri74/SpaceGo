package main

import (
	"io"
	"net"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"
)

var (
	sampleServerPort string
	sampleStartOnce  sync.Once
)

func getSampleFreePort() (string, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return "", err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return "", err
	}
	defer l.Close()
	return ":" + strings.Split(l.Addr().String(), ":")[1], nil
}

func setupSampleServer(t *testing.T) string {
	sampleStartOnce.Do(func() {
		port, err := getSampleFreePort()
		if err != nil {
			t.Fatalf("خطا در یافتن پورت آزاد: %v", err)
		}
		sampleServerPort = port

		go StartServer(port)

		// Wait for server to start
		for i := 0; i < 20; i++ {
			conn, err := net.DialTimeout("tcp", "localhost"+port, 50*time.Millisecond)
			if err == nil {
				conn.Close()
				break
			}
			time.Sleep(50 * time.Millisecond)
		}
	})
	return sampleServerPort
}

func TestPingEndpointSample_ReturnsPong(t *testing.T) {
	port := setupSampleServer(t)

	resp, err := http.Get("http://localhost" + port + "/ping")
	if err != nil {
		t.Fatalf("خطا در برقراری ارتباط با سرور: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("انتظار می‌رفت وضعیت کد ۲۰۰ باشد، اما %d دریافت شد", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("خطا در خواندن پاسخ سرور: %v", err)
	}

	body := string(bodyBytes)
	if body != "pong" {
		t.Errorf("انتظار می‌رفت پاسخ سرور pong باشد، اما '%s' دریافت شد", body)
	}
}
