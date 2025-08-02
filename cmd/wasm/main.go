package main

import (
	"bytes"
	"log"
	"strconv"
	"syscall/js"

	"github.com/cowork-ai/go-diceware"
)

func main() {
	s, err := diceware.NewSamplerFromEFFWordlist(bytes.NewReader(diceware.EFFLargeWordlist))
	if err != nil {
		log.Fatal(err)
	}
	js.Global().Set("sampleWords", sampleWordsWrapper(s))
	<-make(chan struct{})
}

type Sampler interface {
	SampleWords(n int) ([]string, error)
}

func sampleWordsWrapper(s Sampler) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		n, err := strconv.Atoi(args[0].String())
		if err != nil {
			return err.Error()
		}
		words, err := s.SampleWords(n)
		if err != nil {
			return err.Error()
		}
		array := make([]interface{}, len(words))
		for i, w := range words {
			array[i] = w
		}
		return js.ValueOf(array)
	})
}
