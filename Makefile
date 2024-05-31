build: 
	go build -o ./bin/laygochain

run: build
	./bin/laygochain

test:
	go test ./... --v
