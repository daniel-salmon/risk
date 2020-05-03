package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/peterbourgon/ff/v3"
)

func main() {
	var (
		port = flag.Int("port", 8080, "Port on which to run Risk backend")
	)
	if err := ff.Parse(flag.CommandLine, os.Args[1:], ff.WithEnvVarNoPrefix()); err != nil {
		log.Fatalf("Error parsing flags: %s", err)
	}

	router := gin.Default()

	// Set up routes
	setUpRoutes(router)

	// Listen and serve
	router.Run(fmt.Sprintf(":%d", *port))
}

func setUpRoutes(router *gin.Engine) {
	// Health check endpoint
	router.GET("/health", healthHandler)

	// Create a new game
	router.POST("/game", newGameHandler)
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, Success{Success: true})
}

func newGameHandler(c *gin.Context) {
	var newGame NewGame
	if err := c.ShouldBindJSON(&newGame); err != nil {
		e := Error{
			Success: false,
			Message: fmt.Sprintf("Missing required fields %q and %q", "name", "players"),
		}
		handleError(c, http.StatusBadRequest, err, e)
		return
	}
}

// handleError logs the internal error encountered by the service
// and returns the provided HTTP status code along with the custom error message to the client
func handleError(c *gin.Context, HTTPStatusCode int, err error, customError Error) {
	c.Error(err)
	c.JSON(HTTPStatusCode, customError)
}
