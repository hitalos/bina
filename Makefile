build: build_public
	CGO_ENABLED=0 go build -ldflags '-s -w' -trimpath -o bina ./cmd

build_linux:
	CGO_ENABLED=0 GOOS=linux go build -ldflags '-s -w' -trimpath -o bina_linux ./cmd

build_windows:
	CGO_ENABLED=0 GOOS=windows go build -ldflags '-s -w' -trimpath -o bina.exe ./cmd

build_macosx:
	CGO_ENABLED=0 GOOS=darwin go build -ldflags '-s -w' -trimpath -o bina_macosx ./cmd

build_android:
	GOOS=linux GOARCH=arm GOARM=7 go build -o bina_android ./cmd

all: build_linux build_windows build_macosx

lint:
	golangci-lint run ./...

build_public: install_deps
	node ./build.js

test:
	npm test

install_deps:
	npm i

container_image:
	docker build -t bina .

clean:
	rm -rf bina* node_modules cmd/public/scripts cmd/public/styles
