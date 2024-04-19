package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Result struct {
	Title             string  `json:"title"`
	Price             float64 `json:"price"`
	AvailableQuantity int     `json:"available_quantity"`
}

type Response struct {
	SiteID  string   `json:"site_id"`
	Results []Result `json:"results"`
}

func main() {
	fmt.Println("Busca un producto: ")
	input := readInput()
	results := search(input)
	for i, result := range results {
		// []
		fmt.Println(fmt.Sprintf(" [%d] %s ", i, result.Title))
		time.Sleep(200 * time.Millisecond)
	}
}

func readInput() string {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error reading input")
	}
	return input
}

func search(query string) []Result {
	url := fmt.Sprintf("https://api.mercadolibre.com/sites/MLA/search?q=%s", query)
	response, err := http.Get(url)
	if err != nil {
		return nil
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil
	}
	var searchResponse Response
	err = json.Unmarshal(bytes, &searchResponse)
	if err != nil {
		return nil
	}
	return searchResponse.Results
}
