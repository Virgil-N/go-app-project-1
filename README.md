#

## INFO

### wasm 编译

- Linux和Mac系统上编译wasm文件，命令为：`GOARCH=wasm GOOS=js go build -o web/app.wasm`
- Windows系统上为：`$env:GOARCH="wasm";$env:GOOS="js"; go build -o web/app.wasm`

### 运行

- 每次运行程序前必须先编译 wasm 文件
