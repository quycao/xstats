package util

import "strings"

// WordWrap wrap word in console
func WordWrap(text string, lineWidth int) string {
	words := strings.Fields(strings.TrimSpace(text))
	// words := strings.Split(strings.TrimSpace(text), " ")
	if len(words) == 0 {
		return text
	}
	wrapped := words[0]
	spaceLeft := lineWidth - len(wrapped)
	for _, word := range words[1:] {
		// word = strings.TrimSpace(word)
		if len(word)+1 > spaceLeft {
			wrapped += "\n" + word
			spaceLeft = lineWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}

	return wrapped
}
