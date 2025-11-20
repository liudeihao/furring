build:
	go build -o bin/app.exe

run:
	go run server/cmd/api/main.go

test:
	go test .\... -count=1