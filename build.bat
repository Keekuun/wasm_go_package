set GOOS=js
set GOARCH=wasm
go build -o test.wasm test/main.go