build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build -o ./app

run: build

# GOARCH=wasm GOOS=js go build -o web/app.wasm