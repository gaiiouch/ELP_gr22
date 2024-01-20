package main

import (
	"fmt"
	"os"
)

func EcritureMat(taille int, mat [taille][taille]int, filename string) error {
	/*
		écrit une matrice dans un fichier à partir d'une variable matrice déjà créée
	*/

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close() // assurer la fermeture du fichier si le programme plante

	for i := 0; i < taille; i++ {
		for j := 0; j < taille; j++ {
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
