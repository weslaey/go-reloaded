package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1. Lire sample.txt
	content, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("Erreur de lecture :", err)
		return
	}

	text := string(content)
	fmt.Println("Texte lu :", text) // Pour voir ce qui est lu dans le terminal

	// Remplacement simple direct
	// Si le texte contient "(cap)", on met une majuscule au mot juste avant "bridge"
	if strings.Contains(text, "bridge (cap)") {
		text = strings.ReplaceAll(text, "bridge (cap)", "Bridge")
	}

	// 2. Écrire dans result.txt
	err = os.WriteFile("result.txt", []byte(text), 0644)
	if err != nil {
		fmt.Println("Erreur d'écriture :", err)
		return
	}

	fmt.Println("C'est fait ! Regarde result.txt")
}