vet:
	go fmt
	go vet

tidy:
	go fmt
	go mod tidy

parse: 
	go run main.go

test next-token:
	go test -v ./lexer/
