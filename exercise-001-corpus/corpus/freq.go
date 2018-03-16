package corpus

type WordFreq struct {
	Str   string
	Count int
}

// converts a []string to a map[string]int containing
// counts of each string in words
func AsFrequencyMap(words []string) map[string]int {
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	return counts
}

// converts a map[string]int to a []WordFreq
func AsWordFreqs(counts map[string]int) []WordFreq {
	wordFreqs := make([]WordFreq, 0, len(counts))
	for word, count := range counts {
		wordFreqs = append(wordFreqs, WordFreq{word, count})
	}
	return wordFreqs
}

// converts a []string directly to a []WordFreq
func GenerateWordFreqs(words []string) []WordFreq {
	return AsWordFreqs(AsFrequencyMap(words))
}
