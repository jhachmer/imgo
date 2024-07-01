cli:
	go build -o bin/imgo cmd/imgo/main.go

server:
	go build -o bin/imgo-server cmd/imgo_server/main.go

run:
	go run main.go