package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)



// Met une majuscule à la première lettre d'un mot
func capitalize(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// Traite toutes les balises entre parenthèses 
func processTags(words []string) []string {
	for i := 0; i < len(words); i++ {
		word := words[i]

		// Cas des balises simples
		if (word == "(up)" || word == "(low)" || word == "(cap)" || word == "(hex)" || word == "(bin)") && i > 0 {
			switch word {
			case "(up)":
				words[i-1] = strings.ToUpper(words[i-1])
			case "(low)":
				words[i-1] = strings.ToLower(words[i-1])
			case "(cap)":
				words[i-1] = capitalize(words[i-1])
			case "(hex)":
				if val, err := strconv.ParseInt(words[i-1], 16, 64); err == nil {
					words[i-1] = strconv.FormatInt(val, 10)
				}
			case "(bin)":
				if val, err := strconv.ParseInt(words[i-1], 2, 64); err == nil {
					words[i-1] = strconv.FormatInt(val, 10)
				}
			}
			// Supprimer la balise de la liste
			words = append(words[:i], words[i+1:]...)
			i--

			// Cas des balises avec paramètre numérique
		} else if strings.HasPrefix(word, "(") && strings.HasSuffix(word, ")") {
			contentTag := word[1 : len(word)-1]
			parts := strings.Split(contentTag, ",")
			if len(parts) == 2 {
				action := strings.TrimSpace(parts[0])
				n, err := strconv.Atoi(strings.TrimSpace(parts[1]))
				if err == nil && n > 0 && i >= n {
					for j := i - n; j < i; j++ {
						switch action {
						case "up":
							words[j] = strings.ToUpper(words[j])
						case "low":
							words[j] = strings.ToLower(words[j])
						case "cap":
							words[j] = capitalize(words[j])
						}
					}
					// Supprimer la balise de la liste
					words = append(words[:i], words[i+1:]...)
					i--
				}
			}
		}
	}
	return words
}

// Change "a" en "an" si le mot suivant commence par une voyelle ou un h
func processArticles(words []string) []string {
	for i := 0; i < len(words)-1; i++ {
		lowerWord := strings.ToLower(words[i])
		if lowerWord == "a" {
			nextWord := words[i+1]
			if len(nextWord) > 0 {
				firstChar := unicode.ToLower(rune(nextWord[0]))
				if firstChar == 'a' || firstChar == 'e' || firstChar == 'i' || firstChar == 'o' || firstChar == 'u' || firstChar == 'h' {
					if words[i] == "A" {
						words[i] = "An"
					} else {
						words[i] = "an"
					}
				}
			}
		}
	}
	return words
}

// Nettoie et espace correctement la ponctuation
func processPunctuation(text string) string {
	// Supprimer l'espace avant la ponctuation
	rePunctBefore := regexp.MustCompile(`\s+([.,!?:;])`)
	text = rePunctBefore.ReplaceAllString(text, "$1")

	// Ajouter un espace après la ponctuation si elle est suivie d'un caractère non vide
	rePunctAfter := regexp.MustCompile(`([.,!?:;])([^\s])`)
	text = rePunctAfter.ReplaceAllString(text, "$1 $2")

	// Gérer les guillemets simples
	reQuotesOpen := regexp.MustCompile(`'\s+`)
	text = reQuotesOpen.ReplaceAllString(text, "'")
	reQuotesClose := regexp.MustCompile(`\s+'`)
	text = reQuotesClose.ReplaceAllString(text, "'")

	return text
}


func main() {
	// Lire le fichier d'entrée
	content, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("Erreur de lecture :", err)
		return
	}

	text := string(content)
	words := strings.Fields(text)

	// Appliquer les différentes étapes de transformation
	words = processTags(words)
	words = processArticles(words)

	// Regrouper les mots en une seule chaîne de caractères
	processedText := strings.Join(words, " ")

	// Appliquer le nettoyage de la ponctuation sur le texte final
	processedText = processPunctuation(processedText)

	// Écrire le résultat dans le fichier de sortie
	err = os.WriteFile("result.txt", []byte(processedText), 0644)
	if err != nil {
		fmt.Println("Erreur d'écriture :", err)
		return
	}

	fmt.Println("Code exécuté avec succès (version ordonnée) ! Regarde result.txt")
}