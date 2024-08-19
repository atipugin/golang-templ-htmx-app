BINPATH = bin/myapp

.PHONY: build
build: build-templ build-app

.PHONY: build-app
build-app:
	go build -o $(BINPATH) cmd/myapp/main.go

.PHONY: build-templ
build-templ:
	templ generate

.PHONY: run
run: build
	$(BINPATH)

.PHONY: fmt
fmt:
	templ fmt internal/view
