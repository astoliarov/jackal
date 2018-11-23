build-dirs:
	@mkdir -p bin/

build: build-dirs
	go build -o bin/api  ./cmd/api
