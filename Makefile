all: dist/ejson.wasm dist/index.js
all:

dist/index.js: main.js
	cat "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" > dist/index.js && \
    cat main.js >> dist/index.js

dist/ejson.wasm: main.go go.sum
	GOOS=js GOARCH=wasm go build -o $@
