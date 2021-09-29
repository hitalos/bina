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
	go get -u golang.org/x/lint/golint
	${GOPATH}/bin/golint

build_public: install_deps
	mkdir -p cmd/public/styles cmd/public/scripts
	npm test && npm run build
	cat node_modules/vue-material/dist/vue-material.css > cmd/public/styles/bundle.css
	echo >> cmd/public/styles/bundle.css
	cat src/app.css >> cmd/public/styles/bundle.css

install_deps:
	npm i

clean:
	rm -rf bina* node_modules cmd/public/scripts cmd/public/styles
