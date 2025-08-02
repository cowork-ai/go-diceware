package main

import (
	"crypto/rand"
	_ "embed"
	"encoding/csv"
	"io"
	"math/big"
)

//go:embed eff_large_wordlist.txt
var effLargeWordlist []byte

type wordSampler struct {
	words     []string
	wordCount *big.Int
}

func NewSampler(words []string) *wordSampler {
	return &wordSampler{
		words:     words,
		wordCount: big.NewInt(int64(len(words))),
	}
}

// NewSamplerFromEFFWordlist takes an EFF word list in a TSV format (a CSV with a tab delimiter).
// Check out https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt
// or https://web.archive.org/web/20230708051203/https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt
func NewSamplerFromEFFWordlist(r io.Reader) (*wordSampler, error) {
	reader := csv.NewReader(r)
	reader.Comma = '\t'
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	words := make([]string, len(records))
	for i, record := range records {
		if len(record) != 2 {
			return nil, csv.ErrFieldCount
		}
		words[i] = record[1]
	}
	return NewSampler(words), nil
}

func (g *wordSampler) SampleWords(n int) ([]string, error) {
	samples := make([]string, n)
	for i := range n {
		index, err := rand.Int(rand.Reader, g.wordCount)
		if err != nil {
			return nil, err
		}
		samples[i] = g.words[index.Int64()]
	}
	return samples, nil
}
