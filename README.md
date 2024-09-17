# Client-Server API

This project consists of a client-server application where the server fetches currency exchange rates and the client retrieves and saves the data to a file.

## Project Structure

- `client.go`: The client application that requests data from the server and saves it to a file.
- `server.go`: The server application that fetches currency exchange rates from an external API and serves it to the client.
- `go.mod`: The Go module file.
- `cotacao.txt`: The file where the client saves the fetched data.
- `Makefile`: Contains commands to start the server, run the client, and list bids from the database.

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/victormilk/client-server-api.git
    cd client-server-api
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

## Usage

### Start the Server

To start the server, has two ways:

  ```sh
  make start-server
  ```

  ```sh
  go run cmd/server/server.go
  ```

### Run the Client

  To run the client and fetch the latest currency exchange rates:

   ```sh
   make run-client
   ```

   ```sh
   go run cmd/client/client.go
   ```

### List Bids

  To list the bids stored in the database:

  ```sh
  make list-bids
  ```

  ```sh
  sqlite3 bids.db "SELECT * FROM bids;"
  ```
