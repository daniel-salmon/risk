package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/daniel-salmon/risk/game"
	"github.com/daniel-salmon/risk/stores"

	"github.com/gin-gonic/gin"
	"github.com/peterbourgon/ff/v3"
)

var (
	store stores.Store
)

func main() {
	var (
		port = flag.Int("port", 8080, "Port on which to run Risk backend")
	)
	if err := ff.Parse(flag.CommandLine, os.Args[1:], ff.WithEnvVarNoPrefix()); err != nil {
		log.Fatalf("Error parsing flags: %s", err)
	}

	// Create game store
	store, err := stores.NewStore()
	if err != nil {
		log.Fatalf("Error building store: %s", err)
	}
	defer store.Close()

	// Create gin HTTP router
	router := gin.Default()

	// register middleware
	registerMiddleware(router)

	// Set up routes
	setUpRoutes(router)

	// Listen and serve
	router.Run(fmt.Sprintf(":%d", *port))
}

func registerMiddleware(router *gin.Engine) {
	router.HandleMethodNotAllowed = true

	// Require Accept: application/json in Header
	router.Use(func(c *gin.Context) {
		// NOTE: It's entirely possible that more than one Accept header has been specified and this
		// will miss the one that is 'application/json', but this should be an edge case and I don't
		// mind rejecting those requests
		err := errors.New("Request's HTTP 'Accept' header does not match 'application/json'")
		if c.Request.Header.Get("Accept") != "application/json" {
			c.AbortWithError(http.StatusNotAcceptable, err)
		}
	})

	// Require *only* Content-Type: application/json in Header
	router.Use(func(c *gin.Context) {
		err := errors.New("Request's HTTP 'Content-Type' header is invalid, requires *only* 'application/json'")
		if c.Request.Header.Get("Content-Type") != "application/json" {
			c.AbortWithError(http.StatusUnsupportedMediaType, err)
		}
	})
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
		e := &Error{
			Success: false,
			Message: fmt.Sprintf("Missing required fields %q and %q", "name", "players"),
		}
		handleError(c, http.StatusBadRequest, err, e)
		return
	}

	g, err := store.CreateGame(newGame.Name, newGame.Players)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err, nil)
		return
	}

	// Transform the game object into the game response object
	// This removes any data stored in the keys of the game object
	gameResponse := GameResponse{
		Name:          g.Name,
		GoldenCavalry: g.GoldenCavalry,
		Players:       g.Players,
		Territories:   []TerritoryResponse{},
	}

	// Build the territories response object
	for _, territory := range g.Territories {
		t := TerritoryResponse{
			Name:      territory.Name,
			Continent: territory.Continent,
			Links:     territory.Links,
			OwnedBy:   territory.OwnedBy,
			Armies:    []ArmyResponse{},
		}
		for army, value := range territory.Armies {
			a := ArmyResponse{
				Type:  army.String(),
				Value: value,
			}
			t.Armies = append(t.Armies, a)
		}

		gameResponse.Territories = append(gameResponse.Territories, t)
	}

	// Build the cards response object
	cardsResponse := CardsResponse{
		DrawPile:    (*g.Cards).DrawPile,
		DiscardPile: (*g.Cards).DiscardPile,
		Owned:       []CardOwnedResponse{},
	}

	// For convenience we break out the slice of players into a map keyed by the player id
	// This makes it easier to unwrap the cards objects owned by players
	pMap := make(map[int](game.Player))
	for _, p := range g.Players {
		pMap[p.ID] = p
	}

	for pID, cards := range (*g.Cards).OwnedBy {
		for _, card := range cards {
			k := CardOwnedResponse{
				OwnedBy: pMap[pID],
				Card:    card,
			}
			cardsResponse.Owned = append(cardsResponse.Owned, k)
		}
	}

	// Add the cards response to the game reponse object
	gameResponse.Cards = cardsResponse

	c.JSON(http.StatusOK, gameResponse)
}

// handleError logs the internal error encountered by the service
// and returns the provided HTTP status code along with the custom error message to the client
func handleError(c *gin.Context, HTTPStatusCode int, err error, customError *Error) {
	// Write the error to gin's context
	c.Error(err)

	// For internal server errors, we always send a default response message to the user
	// Otherwise we use whatever custom error we've passed to the function
	switch HTTPStatusCode {
	case http.StatusInternalServerError:
		e := Error{Success: false, Message: "Oops! Something unexpected happened"}
		c.JSON(http.StatusInternalServerError, e)
	default:
		c.JSON(HTTPStatusCode, *customError)
	}
}
