package corpus

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCorpusGrepWords(t *testing.T) {
	words := GrepWords("&a&$a $bbb b ...bb.bb.c'c^c'c c#b")
	assert.Equal(t, len(words), 10)

	expected := []string{`a`, `a`, `bbb`, `b`, `bb`, `bb`, `c'c`, `c'c`, `c`, `b`}
	for idx, word := range expected {
		assert.Equal(t, words[idx], word)
	}
}
