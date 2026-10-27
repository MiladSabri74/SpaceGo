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

    // Create a new router and register the grouped routes
    r := gin.Default()
    api := r.Group("/api")
    v1 := api.Group("/v1")
    public := v1.Group("/public")
    public.GET("/songs", handlers.GetSongs)

    // Send a request to the grouped route
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/api/v1/public/songs", nil)
    r.ServeHTTP(w, req)

    // Verify the status code is 200 OK
    if w.Code != http.StatusOK {
        t.Errorf("مسیر گروه‌بندی شده باید وضعیت %d برگرداند، اما %d دریافت شد.", http.StatusOK, w.Code)
    }
}
