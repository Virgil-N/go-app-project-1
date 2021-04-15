build:
	GOARCH=wasm GOOS=js go build -o app/web/app.wasm
	go build -o ./app/app

run: build