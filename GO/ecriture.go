package main

import (
	"fmt"
	"os"
)

// écrit une matrice dans un fichier à partir d'une variable matrice déjà créée
func EcritureMat(taille int, mat [taille]string, filename string) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer file.Close() // assurer la fermeture du fichier si le programme plante

	for i := 0; i < taille; i++ {
		_, err = fmt.Fprintf(file, "%s ", mat[i])
		if err != nil {
			return err

		}

		_, err = fmt.Fprintf(file, "\n")
		if err != nil {
			return err
		}
	}
	return err
}
