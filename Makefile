run:
	go run .

debug:
	DEBUG=1 go run .

fmt:
	gofmt -s -w .

check:
	gocritic check ./...
