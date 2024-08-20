BINPATH = bin/myapp

.PHONY: build
build: build-templ build-css build-app

.PHONY: build-app
build-app:
	go build -o $(BINPATH) cmd/myapp/main.go

.PHONY: build-templ
build-templ:
	templ generate

.PHONY: build-css
build-css:
	npm --prefix web run build:css -- --minify

.PHONY: watch
watch:
	$(MAKE) -j3 watch-app watch-templ watch-css

.PHONY: watch-app
watch-app:
	go run github.com/air-verse/air@latest \
	--build.cmd "$(MAKE) build-app" \
	--build.bin "$(BINPATH)" \
	--build.include_ext "go" \
	--build.exclude_dir "bin,web"

.PHONY: watch-templ
watch-templ:
	templ generate \
	--watch \
	--proxy "http://localhost:8080" \
	--open-browser=false

.PHONY: watch-css
watch-css:
	npm --prefix web run build:css -- --watch=always

.PHONY: fmt
fmt:
	templ fmt internal/view
