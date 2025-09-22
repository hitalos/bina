build: build_public
	CGO_ENABLED=0 go build -ldflags '-s -w' -trimpath -o dist/bina ./cmd

build_linux:
	CGO_ENABLED=0 GOOS=linux go build -ldflags '-s -w' -trimpath -o dist/bina_linux ./cmd

build_windows:
	CGO_ENABLED=0 GOOS=windows go build -ldflags '-s -w' -trimpath -o dist/bina.exe ./cmd

build_macosx:
	CGO_ENABLED=0 GOOS=darwin go build -ldflags '-s -w' -trimpath -o dist/bina_macosx ./cmd

build_android:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -ldflags '-s -w' -trimpath -o dist/bina_android ./cmd

all: build_linux build_windows build_macosx

dev:
	go run -tags dev ./cmd

lint:
	golangci-lint run ./...

sec:
	go vet ./...
	govulncheck ./...
	grype .
	trivy fs .

build_public: install_deps js css

js:
	npm run build::js

css:
	npm run build::css

install_deps:
	npm ci

IMAGE_BUILDER=$(shell [ -e /usr/bin/buildah ] && echo buildah || echo docker)
CONTAINER_REGISTRY?=localhost
container_image:
	$(IMAGE_BUILDER) build -t $(CONTAINER_REGISTRY)/nti/bina:latest .

container_image_sec:
	trivy image $(CONTAINER_REGISTRY)/nti/bina:latest
	grype $(CONTAINER_REGISTRY)/nti/bina:latest

container_image_push:
	$(IMAGE_BUILDER) push $(CONTAINER_REGISTRY)/nti/bina:latest

clean:
	rm -rf dist node_modules cmd/public/assets/scripts/*.min.js* cmd/public/assets/styles/*.min.css*
