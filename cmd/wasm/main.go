package main

import (
	"bytes"
	"strconv"
	"syscall/js"

	"github.com/cowork-ai/go-diceware"
)

type Sampler interface {
	SampleWords(n int) ([]string, error)
}

var defaultSampler = func() Sampler {
	s, err := diceware.NewSamplerFromEFFWordlist(bytes.NewReader(diceware.EFFLargeWordlist))
	if err != nil {
		panic(err)
	}
	return s
}()

func main() {
	js.Global().Set("sampleWords", js.FuncOf(sampleWords))
	select {}
}

func sampleWords(this js.Value, args []js.Value) interface{} {
	n, err := strconv.Atoi(args[0].String())
	if err != nil {
		panic(err)
	}
	words, err := defaultSampler.SampleWords(n)
	if err != nil {
		panic(err)
	}
	result := make([]interface{}, len(words))
	for i, word := range words {
		result[i] = word
	}
	return js.ValueOf(result)
}
