APP_NAME = go-pwa

build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build

run: build
	./$(APP_NAME)

tailwind:
	npx tailwindcss -i ./web/styles.css -o ./web/tailwind.css -w

dev:
	air
