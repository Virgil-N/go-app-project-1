build:
	chmod a+x ./bin; GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build -o ./bin

run: build