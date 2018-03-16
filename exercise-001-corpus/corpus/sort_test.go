package corpus

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCorpusSortWordFreqs(t *testing.T) {
	freqs := []WordFreq{
		WordFreq{`they`, 5},
		WordFreq{`our`, 4},
		WordFreq{`his`, 4},
		WordFreq{`be`, 2},
		WordFreq{`am`, 2},
		WordFreq{`he`, 2},
		WordFreq{`a`, 1},
		WordFreq{`i`, 1},
		WordFreq{`us`, 1},
	}
	SortWordFreqs(freqs)
	expected := []string{`they`, `our`, `his`, `be`, `he`, `am`, `i`, `us`, `a`}
	for i, word := range expected {
		assert.Equal(t, freqs[i].Str, word)
	}
}
