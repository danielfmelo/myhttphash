img := danielfmelo/myhttphash:latest
wd := $(shell pwd)
cachevol=$(wd):/go/pkg/mod
rundocker := docker run --rm -v $(wd):/app -v $(cachevol) $(img)
urls?=
parallel?=10

image:
	docker build . -t $(img)

run: 
	go run cmd/myhttphash.go -parallel $(parallel) $(urls)

docker-run: image
	$(rundocker) go run cmd/myhttphash.go -parallel $(parallel) $(urls)

build:
	go build -o ./myhttphash ./cmd/myhttphash.go

docker-build: image
	$(rundocker) go build -v -o ./myhttphash ./cmd/myhttphash.go

tests: 
	go test -timeout 20s -tags unit -race -coverprofile=coverage.out ./...

docker-tests: image
	$(rundocker) go test -timeout 20s -tags unit -race -coverprofile=coverage.out ./...

coverage: tests
	go tool cover -html=coverage.out -o=coverage.html
	xdg-open coverage.html
