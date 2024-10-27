build:
	@go build -o bin/qrgo

run: build
	@./bin/qrgo
