APP_NAME = website-wasm

build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build -o $(APP_NAME)

run: build
	./$(APP_NAME) local

build-github: clean
	GOARCH=wasm GOOS=js go build -o docs/web/app.wasm
	go build -o $(APP_NAME)
	@cp web/*.css docs/web/
	./$(APP_NAME) github

tailwind:
	npx tailwindcss -i ./web/styles.css -o ./web/tailwind.css -w

dev:
	air

clean:
	@go clean ./...
