#

## INFO

### 编译

- 在Linux和Mac系统上编译wasm文件，命令为：`GOARCH=wasm GOOS=js go build -o web/app.wasm`，Windows系统上为：`$env:GOARCH="wasm";$env:GOOS="js"; go build -o web/app.wasm`
