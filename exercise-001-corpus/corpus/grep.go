package corpus

import (
	"regexp"
	"strings"
)

// GrepWords ingests a string and outputs an array
// of 'words' in that string, which are sequences
// of alphanumeric characters and single apostrophes
func GrepWords(str string) []string {
	if len(str) == 0 {
		return []string{}
	}

	str = strings.ToLower(str)
	re := regexp.MustCompile(`\b[\w\']+\b`)
	return re.FindAllString(str, -1)
}
