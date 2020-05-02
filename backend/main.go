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

	// Create router
	router := gin.Default()

	// Set up ping endpoint
	router.GET("/ping", ping)

	// Listen and serve
	router.Run(fmt.Sprintf(":%d", *port))
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Success": true})
}
