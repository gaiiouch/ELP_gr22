package main

import (
	"fmt"
	"os"
)

func EcrireMatriceDansFichier(n int, matrice [][]int, nomFichier string) error {

	fichier, err := os.Create(nomFichier)
	if err != nil {
		return err
	}
	defer fichier.Close()

	// ecriture dans le fichier
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			_, err = fmt.Fprintf(fichier, "%d ", matrice[i][j])
			if err != nil {
				return err
			}
		}
		_, err = fmt.Fprintf(fichier, "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
