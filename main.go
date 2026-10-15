package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Fonction pour mettre une majuscule initiale
func capitalize(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func main() {
	// 1. Lire le fichier d'entrée
	content, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("Erreur de lecture :", err)
		return
	}

	text := string(content)
	words := strings.Fields(text)

	// 2. Parcourir les mots pour chercher les balises
	for i := 0; i < len(words); i++ {
		word := words[i]

		// Cas des balises simples : (up), (low), (cap)
		if (word == "(up)" || word == "(low)" || word == "(cap)") && i > 0 {
			if word == "(up)" {
				words[i-1] = strings.ToUpper(words[i-1])
			} else if word == "(low)" {
				words[i-1] = strings.ToLower(words[i-1])
			} else if word == "(cap)" {
				words[i-1] = capitalize(words[i-1])
			}
			// Supprimer la balise
			words = append(words[:i], words[i+1:]...)
			i--
		} else if strings.HasPrefix(word, "(") && strings.HasSuffix(word, ")") {
			// Cas des balises avec un nombre, ex: (up, 2)
			contentTag := word[1 : len(word)-1] // Enlève les parenthèses -> "up, 2"
			parts := strings.Split(contentTag, ",")
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

	processedText := strings.Join(words, " ")

	// 3. Écrire le résultat
	err = os.WriteFile("result.txt", []byte(processedText), 0644)
	if err != nil {
		fmt.Println("Erreur d'écriture :", err)
		return
	}

	fmt.Println("Traité avec succès ! Regarde result.txt")
}