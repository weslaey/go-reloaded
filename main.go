package main

import (
	"fmt"
	"os"
)

func main() {
	// Vérifier la présence des arguments 
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Usage: go run . <inputFile> <outputFile>")
		return
	}

	inputFile := args[0]
	outputFile := args[1]

	// Lire le fichier d'entrée
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier : %v\n", err)
		return
	}

	// Pour l'instant, on prépare le texte
	processedText := string(content)

	// Écrire dans le fichier de sortie
	err = os.WriteFile(outputFile, []byte(processedText), 0644)
	if err != nil {
		fmt.Printf("Erreur lors de l'écriture du fichier : %v\n", err)
		return
	}

	fmt.Println("Fichier traité avec succès !")
}