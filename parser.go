package main

import (
	"strconv"
	"strings"
	"unicode"
)

// Fonction utilitaire pour mettre une majuscule initiale (Cap)
func capitalize(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	// Optionnel : mettre le reste en minuscules si besoin, 
	// mais généralement le sujet demande juste la majuscule initiale.
	return string(runes)
}

// ProcessText gère toutes les transformations
func ProcessText(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		word := words[i]

		// Gestion des balises simples : (up), (low), (cap)
		if (word == "(up)" || word == "(low)" || word == "(cap)") && i > 0 {
			if word == "(up)" {
				words[i-1] = strings.ToUpper(words[i-1])
			} else if word == "(low)" {
				words[i-1] = strings.ToLower(words[i-1])
			} else if word == "(cap)" {
				words[i-1] = capitalize(words[i-1])
			}
			// Supprimer la balise du tableau
			words = append(words[:i], words[i+1:]...)
			i--
		} else if strings.HasPrefix(word, "(") && strings.HasSuffix(word, ")") {
			// Gestion des balises avec un nombre, ex: (up, 2)
			content := word[1 : len(word)-1] // Enlève les parenthèses -> "up, 2"
			parts := strings.Split(content, ",")
			if len(parts) == 2 {
				action := strings.TrimSpace(parts[0])
				n, err := strconv.Atoi(strings.TrimSpace(parts[1]))
				if err == nil && n > 0 && i >= n {
					// Appliquer l'action sur les n mots précédents
					for j := i - n; j < i; j++ {
						if action == "up" {
							words[j] = strings.ToUpper(words[j])
						} else if action == "low" {
							words[j] = strings.ToLower(words[j])
						} else if action == "cap" {
							words[j] = capitalize(words[j])
						}
					}
					// Supprimer la balise
					words = append(words[:i], words[i+1:]...)
					i--
				}
			}
		}
	}

	return strings.Join(words, " ")
}