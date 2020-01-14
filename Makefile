build:
	CGO_ENABLED=0 go build -ldflags '-s -w' -trimpath -o bina ./cmd

build_linux:
	CGO_ENABLED=0 GOOS=linux go build -ldflags '-s -w' -trimpath -o bina_linux ./cmd

build_windows:
	CGO_ENABLED=0 GOOS=windows go build -ldflags '-s -w' -trimpath -o bina.exe ./cmd

build_macosx:
	CGO_ENABLED=0 GOOS=darwin go build -ldflags '-s -w' -trimpath -o bina_macosx ./cmd

build_android:
	GOOS=linux GOARCH=arm GOARM=7 go build -o bina_android ./cmd

all: go_generate build_linux build_windows build_macosx

go_generate: build_public
	${GOPATH}/bin/rice -i ./cmd embed-go

build_public: install_deps
	mkdir -p public/styles public/scripts
	npm run build
	cat node_modules/vue-material/dist/vue-material.css > public/styles/bundle.css
	echo >> public/styles/bundle.css
	cat src/app.css >> public/styles/bundle.css

install_go_deps:
	go get -u github.com/GeertJohan/go.rice
	go get -u github.com/GeertJohan/go.rice/rice

install_js_deps:
	npm i

install_deps: install_go_deps install_js_deps

clean:
	rm -rf node_modules package-lock.json public/scripts public/styles cmd/rice-box.go
