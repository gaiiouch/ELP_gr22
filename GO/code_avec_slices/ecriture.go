package main

import (
	"fmt"
	"os"
)

// écrit une matrice dans un fichier à partir d'une variable matrice déjà créée
func EcritureMat(n int, mat [][]int, filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// ecriture dans le fichier
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			_, err = fmt.Fprintf(file, "%d ", mat[i][j])
			if err != nil {
				return err
			}
		}
		_, err = fmt.Fprintf(file, "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
