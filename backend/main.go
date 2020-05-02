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
	router.GET("/health", health)
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, Success{Success: true})
}
