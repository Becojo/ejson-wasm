package main

import (
	"github.com/Shopify/ejson"
	"strings"
	"syscall/js"
)

func ejsonEncrypt(input string) (string, error) {
	inputBuf := strings.NewReader(input)
	outputBuf := new(strings.Builder)

	_, err := ejson.Encrypt(inputBuf, outputBuf)

	return outputBuf.String(), err
}

func ejsonWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return false
		}

		input := args[0].String()
		output, err := ejsonEncrypt(input)

		if err != nil {
			return false
		}

		return output
	})
}

func main() {
	js.Global().Set("ejson", map[string]interface{}{
		"encrypt": ejsonWrapper(),
	})

	select {}
}
