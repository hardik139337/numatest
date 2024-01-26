.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./dist/even-odd-example main.go

.PHONY: image
image: build
	docker build -t "even-odd:v0.5.0" --target even-odd .

clean:
	-rm -rf ./dist
