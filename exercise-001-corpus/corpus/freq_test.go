package corpus

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCorpusAsFrequencyMap(t *testing.T) {
	strs := []string{`a`, `b`, `b`, `c`, `c`, `c`}
	actual := AsFrequencyMap(strs)

	expected := map[string]int{
		`a`: 1,
		`b`: 2,
		`c`: 3,
		`d`: 0,
	}
	for key, val := range expected {
		assert.Equal(t, actual[key], val)
	}
}

func TestCorpusAsWordFreqs(t *testing.T) {
	counts := map[string]int{
		`a`: 1,
		`b`: 2,
		`c`: 3,
	}
	actual := AsWordFreqs(counts)
	for _, freq := range actual {
		assert.Equal(t, freq.Count, counts[freq.Str])
	}
}
