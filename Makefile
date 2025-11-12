build:
	go build -o bin/app.exe

run:
	go run .

test:
	go test .\... -count=1