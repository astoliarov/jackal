build-dirs:
	@mkdir -p bin/

build: build-dirs
	go build -v -o bin/api  ./cmd/api
