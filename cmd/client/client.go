package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type APIResponse struct {
	Bid string `json:"bid"`
}

const (
	API             = "http://localhost:8080/cotacao"
	REQUEST_TIMEOUT = 300 * time.Millisecond
)

func main() {
	ctx := context.Background()

	response, err := makeRequest(ctx)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("cotacao.txt")
	if err != nil {
		log.Fatal(err)
	}

	content := fmt.Sprintf("Dólar: %s\n", response.Bid)
	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Cotação salva com sucesso!")
}

func makeRequest(ctx context.Context) (*APIResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, REQUEST_TIMEOUT)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, API, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response APIResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
