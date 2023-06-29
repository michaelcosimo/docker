package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var JOKE_URL = os.Getenv("JOKE_URL")

type Joke struct {
	Created_at string `json:"created_at"`
	IconUrl    string `json:"icon_url"`
	ID         string `json:"id"`
	UpdatedAt  string `json:"updated_at"`
	Url        string `json:"url"`
	Value      string `json:"value"`
}

type Response struct {
	Joke string `json:"joke"`
}

func main() {
	fmt.Printf("Joke URL: %s\n", JOKE_URL)
	router := gin.Default()
	router.GET("/joke", jokeHandler)
	port := ":8080"
	fmt.Printf("Server listening on port %s\n", port)
	router.Run(port)
}

func jokeHandler(c *gin.Context) {
	resp, err := http.Get(JOKE_URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch joke"})
		return
	}
	defer resp.Body.Close()

	var jokeData Joke
	err = json.NewDecoder(resp.Body).Decode(&jokeData)
	if err != nil {
		fmt.Printf("%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse joke response"})
		return
	}

	res := Response{
		Joke: jokeData.Value,
	}

	c.JSON(http.StatusOK, res)
}
