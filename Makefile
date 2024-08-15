.PHONY: build
build:
	templ generate internal/view
	go build -o bin/myapp cmd/myapp/main.go

.PHONY: run
run: build
	./bin/myapp

.PHONY: fmt
fmt:
	templ fmt internal/view
