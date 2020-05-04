build-wasm:
	GOOS=js GOARCH=wasm go build -o wasm/main.wasm

run-server:
	make build-wasm
	go run server.go