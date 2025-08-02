package diceware

import (
	"bytes"
	"slices"
	"strings"
	"testing"
)

func TestSampleWords(t *testing.T) {
	g, err := NewSamplerFromEFFWordlist(bytes.NewReader(EFFLargeWordlist))
	if err != nil {
		t.Fatal(err)
	}
	words, err := g.SampleWords(3)
	if err != nil {
		t.Fatal(err)
	}
	slices.Sort(words)
	words = slices.Compact(words)
	if got, want := len(words), 3; got != want {
		t.Errorf("len(words)=%v, want=%v", got, want)
	}
	t.Logf("%q\n", words)
	for _, word := range words {
		if got, want := strings.Contains(string(EFFLargeWordlist), word), true; got != want {
			t.Errorf("Contains(word)=%v, want=%v", got, want)
		}
	}
}
