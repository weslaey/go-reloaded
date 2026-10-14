package main

import (
	"fmt"
	"os"
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

	// 2. Parcourir les mots pour chercher (up), (low) ou (cap)
	for i := 0; i < len(words); i++ {
		if i > 0 {
			if words[i] == "(up)" {
				words[i-1] = strings.ToUpper(words[i-1])
				// Supprimer la balise de la liste
				words = append(words[:i], words[i+1:]...)
				i--
			} else if words[i] == "(low)" {
				words[i-1] = strings.ToLower(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			} else if words[i] == "(cap)" {
				words[i-1] = capitalize(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
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