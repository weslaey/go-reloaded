package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func capitalize(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func main() {
	content, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("Erreur de lecture :", err)
		return
	}

	text := string(content)
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		word := words[i]

		// 1. Balises simples : (up), (low), (cap), (hex), (bin)
		if (word == "(up)" || word == "(low)" || word == "(cap)" || word == "(hex)" || word == "(bin)") && i > 0 {
			if word == "(up)" {
				words[i-1] = strings.ToUpper(words[i-1])
			} else if word == "(low)" {
				words[i-1] = strings.ToLower(words[i-1])
			} else if word == "(cap)" {
				words[i-1] = capitalize(words[i-1])
			} else if word == "(hex)" {
				val, err := strconv.ParseInt(words[i-1], 16, 64)
				if err == nil {
					words[i-1] = strconv.FormatInt(val, 10)
				}
			} else if word == "(bin)" {
				val, err := strconv.ParseInt(words[i-1], 2, 64)
				if err == nil {
					words[i-1] = strconv.FormatInt(val, 10)
				}
			}
			words = append(words[:i], words[i+1:]...)
			i--
		} else if strings.HasPrefix(word, "(") && strings.HasSuffix(word, ")") {
			// 2. Balises avec nombre, ex: (up,2)
			contentTag := word[1 : len(word)-1]
			parts := strings.Split(contentTag, ",")
			if len(parts) == 2 {
				action := strings.TrimSpace(parts[0])
				n, err := strconv.Atoi(strings.TrimSpace(parts[1]))
				if err == nil && n > 0 && i >= n {
					for j := i - n; j < i; j++ {
						if action == "up" {
							words[j] = strings.ToUpper(words[j])
						} else if action == "low" {
							words[j] = strings.ToLower(words[j])
						} else if action == "cap" {
							words[j] = capitalize(words[j])
						}
					}
					words = append(words[:i], words[i+1:]...)
					i--
				}
			}
		}
	}

	processedText := strings.Join(words, " ")

	err = os.WriteFile("result.txt", []byte(processedText), 0644)
	if err != nil {
		fmt.Println("Erreur d'écriture :", err)
		return
	}

	fmt.Println("Traité avec succès ! Regarde result.txt")
}