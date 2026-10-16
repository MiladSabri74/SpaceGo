package main

import (
	"GoMusic/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
)

// This is a sample test
func TestGetSongs_Sample(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	// Create a new router and register the route
	r := gin.Default()
	r.GET("/songs", handlers.GetSongs)

	// Send a request to the route
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/songs", nil)
	r.ServeHTTP(w, req)

	// Verify the status code is 200 OK
	if w.Code != http.StatusOK {
		t.Errorf("انتظار می‌رفت وضعیت %d باشد، اما %d دریافت شد.", http.StatusOK, w.Code)
	}
}
