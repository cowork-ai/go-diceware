package main

import (
	"bytes"
	"log"
	"strconv"
	"syscall/js"
)

func main() {
	s, err := NewSamplerFromEFFWordlist(bytes.NewReader(effLargeWordlist))
	if err != nil {
		log.Fatal(err)
	}
	js.Global().Set("sampleWords", sampleWordsWrapper(s))
	<-make(chan struct{})
}

func sampleWordsWrapper(s *wordSampler) js.Func {
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
