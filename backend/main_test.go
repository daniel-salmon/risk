package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
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
	happyHeaders = http.Header{
		"Accept":       []string{"application/json"},
		"Content-Type": []string{"application/json"},
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

func testRequest(method, url string, headers http.Header, body interface{}, statusCode int, expected interface{}, router *gin.Engine, t *testing.T) {
	// Create a new response recorder
	w := httptest.NewRecorder()

	// Marshal the request body
	reqBody, err := json.Marshal(body)
	if err != nil {
		t.Error("Marshaling request body:", err)
		return
	}

	// Create the HTTP request for this test case
	req, err := http.NewRequest(method, url, bytes.NewReader(reqBody))
	if err != nil {
		t.Error("Creating new request:", err)
		return
	}

	// Add the headers for this test request
	for h, v := range headers {
		for _, k := range v {
			req.Header.Add(h, k)
		}
	}

	// Submit the request to the mock server
	router.ServeHTTP(w, req)

	// Read the result
	resp := w.Result()

	// Check the status code
	if resp.StatusCode != statusCode {
		t.Errorf("Expected HTTP Status Code %d, got: %d", statusCode, resp.StatusCode)
		return
	}

	// Return early if there is no expected response body
	// This is usually (but not always!) the case for 400 HTTP status codes
	// where we rely on the status code to convey what went wrong
	if expected == nil {
		return
	}

	// Check the response body
	actualBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("Reading response body:", err)
		return
	}

	// Marshal the expected response to a byte slice
	expectedBytes, err := json.Marshal(expected)
	if err != nil {
		t.Error("Marshaling expected response body:", err)
		return
	}

	// Unmarshal both the actual response and the expected response to JSON
	// and test if the resulting interfaces are deeply equal
	var actualResponse interface{}
	var expectedResponse interface{}
	if err := json.Unmarshal(actualBytes, &actualResponse); err != nil {
		t.Error("Unmarshaling response body:", err)
	} else if json.Unmarshal(expectedBytes, &expectedResponse); err != nil {
		t.Error("Unmarshaling expected response body:", err)
	} else if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Errorf("Expected response \n%s\n got: \n%s\n", expectedResponse, actualResponse)
	}
}

func TestHealth(t *testing.T) {
	testCases := []struct {
		name       string
		method     string
		url        string
		headers    http.Header
		body       interface{}
		statusCode int
		expected   interface{}
	}{
		{
			name:       "InvalidMethod",
			method:     http.MethodPost,
			url:        "/health",
			headers:    happyHeaders,
			body:       nil,
			statusCode: http.StatusMethodNotAllowed,
			expected:   nil,
		},
		{
			name:       "InvalidAcceptHeader",
			method:     http.MethodGet,
			url:        "/health",
			headers:    http.Header{"Accept": []string{"text/plain"}, "Content-Type": []string{"application/json"}},
			body:       nil,
			statusCode: http.StatusNotAcceptable,
			expected:   nil,
		},
		{
			name:       "InvalidContenTypeHeader",
			method:     http.MethodGet,
			url:        "/health",
			headers:    http.Header{"Accept": []string{"application/json"}, "Content-Type": []string{"text/plain"}},
			body:       nil,
			statusCode: http.StatusUnsupportedMediaType,
			expected:   nil,
		},
		{
			name:       "Success",
			method:     http.MethodGet,
			url:        "/health",
			headers:    happyHeaders,
			body:       nil,
			statusCode: http.StatusOK,
			expected:   Success{Success: true},
		},
	}

	router := newMockRouter()
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testRequest(
				testCase.method,
				testCase.url,
				testCase.headers,
				testCase.body,
				testCase.statusCode,
				testCase.expected,
				router,
				t,
			)
		})
	}
}

func TestNewGame(t *testing.T) {
	// TODO: Mock the store so we can generate an HTTP Internal Server Error (500)
	testCases := []struct {
		name       string
		method     string
		url        string
		headers    http.Header
		body       interface{}
		statusCode int
		expected   interface{}
	}{
		{
			name:       "InvalidMethod",
			method:     http.MethodGet,
			url:        "/game",
			headers:    happyHeaders,
			body:       nil,
			statusCode: http.StatusMethodNotAllowed,
			expected:   nil,
		},
		{
			name:       "InvalidAcceptHeader",
			method:     http.MethodPost,
			url:        "/game",
			headers:    http.Header{"Accept": []string{"text/plain"}, "Content-Type": []string{"application/json"}},
			body:       nil,
			statusCode: http.StatusNotAcceptable,
			expected:   nil,
		},
		{
			name:       "InvalidContentTypeHeader",
			method:     http.MethodPost,
			url:        "/game",
			headers:    http.Header{"Accept": []string{"application/json"}, "Content-Type": []string{"text/plain"}},
			body:       nil,
			statusCode: http.StatusUnsupportedMediaType,
			expected:   nil,
		},
		{
			name:       "MissingBody",
			method:     http.MethodPost,
			url:        "/game",
			headers:    happyHeaders,
			body:       nil,
			statusCode: http.StatusBadRequest,
			expected:   Error{Success: false, Message: fmt.Sprintf("Missing required fields %q and %q", "name", "players")},
		},
		{
			name:       "BadRequestBody",
			method:     http.MethodPost,
			url:        "/game",
			headers:    happyHeaders,
			body:       `{"body": "sup"}`,
			statusCode: http.StatusBadRequest,
			expected:   Error{Success: false, Message: fmt.Sprintf("Missing required fields %q and %q", "name", "players")},
		},
		{
			name:       "Success",
			method:     http.MethodPost,
			url:        "/game",
			headers:    happyHeaders,
			body:       newGame,
			statusCode: http.StatusOK,
			expected:   nil,
		},
	}

	router := newMockRouter()
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testRequest(
				testCase.method,
				testCase.url,
				testCase.headers,
				testCase.body,
				testCase.statusCode,
				testCase.expected,
				router,
				t,
			)
		})
	}
}
