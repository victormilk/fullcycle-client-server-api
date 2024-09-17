start-server:
	@echo "Starting server..."
	@go run cmd/server/server.go

run-client:
	@echo "Running client..."
	@go run cmd/client/client.go

list-bids:
	@echo "Listing bids from database..."
	@sqlite3 bids.db "SELECT * FROM bids;"
	@echo "Done."
