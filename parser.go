package main

import (
	"strings"
)

// ProcessText centralise toutes les règles de ton projet
func ProcessText(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		// Règle pour (up) simple
		if words[i] == "(up)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		}
		// Tu pourras rajouter les autres (low), (cap), (hex), etc. ici
	}

	return strings.Join(words, " ")
}