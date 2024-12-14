build:
 @go build -o bin/execute src/main.go

test:
 @go test -v ./...

run: build 
 @./bin/execute