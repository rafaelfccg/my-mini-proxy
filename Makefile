build:
	env GOOS=linux GOARCH=386 go build -a --ldflags="-s" -o bin/new_arbitrage cmd/new_arbitrage/new_arbitrage.go