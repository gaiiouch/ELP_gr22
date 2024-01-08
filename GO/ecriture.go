package main

import (
	"fmt"
	"os"
)

// écrit une matrice dans un fichier à partir d'une variable matrice déjà créée
func EcritureMat(taille int, mat [taille][taille]int, filename string) int {
	file, err := os.Create(filename)

	if err != nil {
		return 1
	}

	for i := 0; i < taille; i++ {
		for j := 0; j < taille; j++ {
			_, err = fmt.Fprintf(file, "%d ", mat[i][j])
			if err != nil {
				return 1
			}
		}
		_, err = fmt.Fprintf(file, "\n")
		if err != nil {
			return 1
		}
	}

	defer file.Close()

	return 0
}
