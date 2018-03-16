package corpus

// WordFreq contains a pairing of a string and
// a count of occurrences of that string
type WordFreq struct {
	Str   string
	Count int
}

// AsFrequencyMap converts a []string to a map[string]int
// containing counts of each string in words
func AsFrequencyMap(words []string) map[string]int {
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	return counts
}

// AsWordFreqs converts a map[string]int to a []WordFreq
func AsWordFreqs(counts map[string]int) []WordFreq {
	wordFreqs := make([]WordFreq, 0, len(counts))
	for word, count := range counts {
		wordFreqs = append(wordFreqs, WordFreq{word, count})
	}
	return wordFreqs
}

// GenerateWordFreqs converts a []string directly to a []WordFreq
func GenerateWordFreqs(words []string) []WordFreq {
	return AsWordFreqs(AsFrequencyMap(words))
}
