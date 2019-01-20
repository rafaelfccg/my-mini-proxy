build:
	env GOOS=linux GOARCH=386 go build -a --ldflags="-s" -o bin/proxy my-simple-proxy.go