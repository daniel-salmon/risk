package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func newMockRouter() *gin.Engine {
	// We set gin to release mode to suppress extra logging that occurs during debug mode
	gin.SetMode(gin.ReleaseMode)

	// We omit any middleware to suppress any extra logging
	// and additional functionality we don't want to test
	router := gin.New()

	// Set up routes
	setUpRoutes(router)

	return router
}

func TestHealth(t *testing.T) {
	router := newMockRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected HTTP Status OK, got: %d", w.Code)
	}
}
