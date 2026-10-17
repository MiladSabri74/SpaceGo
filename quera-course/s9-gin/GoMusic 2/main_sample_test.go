package main

import (
	"GoMusic/handlers"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// This is a sample test
func TestCreateSong_Sample(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new router and register the route
	r := gin.Default()
	r.POST("/songs", handlers.CreateSong)

	// Send an empty request which should fail validation
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/songs", bytes.NewBuffer([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	// Verify the status code is 400 Bad Request
	if w.Code != http.StatusBadRequest {
		t.Errorf("برای دیتای خالی انتظار می‌رفت وضعیت %d باشد، اما %d دریافت شد.", http.StatusBadRequest, w.Code)
	}
}
