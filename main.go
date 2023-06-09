package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const JOKE_URL = "https://api.chucknorris.io/jokes/random"

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
	http.HandleFunc("/joke", jokeHandler)
	port := ":8080"
	fmt.Printf("Server listening on port %s\n", port)
	http.ListenAndServe(port, nil)
}

func jokeHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(JOKE_URL)
	if err != nil {
		http.Error(w, "Failed to fetch joke", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var jokeData Joke
	err = json.NewDecoder(resp.Body).Decode(&jokeData)
	if err != nil {
		fmt.Printf("%v", err)
		http.Error(w, "Failed to parse joke response", http.StatusInternalServerError)
		return
	}

	res := Response{
		Joke: jokeData.Value,
	}

	responseJSON, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("%v", err)
		http.Error(w, "Failed to encode joke response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
