package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	API             = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	REQUEST_TIMEOUT = 200 * time.Millisecond
	PERSIST_TIMEOUT = 10 * time.Millisecond
)

func main() {
	db, err := sql.Open("sqlite3", "bids.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS bids (bid TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("GET /cotacao", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		response, err := makeRequest(ctx)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bid := response.USDBRL.Bid
		err = persistBid(ctx, db, bid)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(APIResponse{Bid: bid})
	})

	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

type AwesomeAPIResponse struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

type APIResponse struct {
	Bid string `json:"bid"`
}

func makeRequest(ctx context.Context) (*AwesomeAPIResponse, error) {
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

	var apiResponse AwesomeAPIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

func persistBid(ctx context.Context, db *sql.DB, bid string) error {
	ctx, cancel := context.WithTimeout(ctx, PERSIST_TIMEOUT)
	defer cancel()

	_, err := db.ExecContext(ctx, "INSERT INTO bids (bid) VALUES (?)", bid)

	return err
}
