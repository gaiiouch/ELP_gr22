package main

import (
	"fmt"
	"os"
)

const taille int = 3

func ecritureMat(mat [taille][taille]int, filename string) int {
	file, err := os.Create(filename)

	if err != nil {
		return 1
	}

	defer file.Close()

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

	return 0
}

func main() {
	var matA [taille][taille]int

	for i := 0; i < taille; i++ {
		for j := 0; j < taille; j++ {
			matA[i][j] = 2
		}
	}

	err := ecritureMat(matA, "matriceA.txt")
	if err == 1 {
		fmt.Println("Erreur lors de l'écriture")
	}
}