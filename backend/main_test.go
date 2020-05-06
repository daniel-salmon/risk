package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daniel-salmon/risk/game"

	"github.com/gin-gonic/gin"
)

var (
	newGame = NewGame{
		Name: "World Domination",
		Players: []game.Player{
			game.Player{ID: 0, Name: "Zero"},
			game.Player{ID: 1, Name: "One"},
			game.Player{ID: 2, Name: "Two"},
		},
	}
)

func newMockRouter() *gin.Engine {
	// We set gin to release mode to suppress extra logging that occurs during debug mode
	gin.SetMode(gin.ReleaseMode)

	// We omit any middleware to suppress any extra logging
	// and additional functionality we don't want to test
	router := gin.New()

	// Register middleware
	registerMiddleware(router)

	// Set up routes
	setUpRoutes(router)

	return router
}

func TestHealth(t *testing.T) {
	router := newMockRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Error("Creating new request: ", err)
	}
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	resp := w.Result()
	fmt.Println(resp.Header)

	if w.Code != http.StatusOK {
		t.Errorf("Expected HTTP Status OK, got: %d", w.Code)
	}

	var success Success
	if err := json.Unmarshal(w.Body.Bytes(), &success); err != nil {
		t.Error("Unmarshalling health check success response: ", err)
	}
}

func TestNewGame(t *testing.T) {
	router := newMockRouter()
	w := httptest.NewRecorder()

	body, err := json.Marshal(newGame)
	if err != nil {
		t.Error("Unmarshalling game1: ", err)
	}
	req, err := http.NewRequest("POST", "/game", bytes.NewReader(body))
	if err != nil {
		t.Error("Creating new request: ", err)
	}
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected HTTP Status OK, got: %d", w.Code)
	}

	var response GameResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Error("Unmarshaling POST /game response: ", err)
	}
}
