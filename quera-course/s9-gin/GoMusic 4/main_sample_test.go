package main

import (
    "GoMusic/middlewares"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
)

// This is a sample test
func TestCustomRecovery_Sample(t *testing.T) {
    gin.SetMode(gin.TestMode)

    // Create a new router and apply the middleware
    r := gin.New()
    r.Use(middlewares.CustomRecovery())

    // Add a route that panics
    r.GET("/panic", func(c *gin.Context) {
        panic("test panic")
    })

    // Send a request to the panic route
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/panic", nil)
    r.ServeHTTP(w, req)

    // Verify the CustomRecovery caught it and returned 500
    if w.Code != http.StatusInternalServerError {
        t.Errorf("میدلور ریکاوری باید ارور ۵۰۰ برگرداند، اما %d دریافت شد.", w.Code)
    }
}
