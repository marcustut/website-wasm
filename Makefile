APP_NAME = go-pwa

build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build

run: build
	./$(APP_NAME)

dev:
	nodemon -e go,css --exec make run --signal SIGTERM