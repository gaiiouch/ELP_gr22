package main

import (
	"fmt"
	"log"
	"os"
)

// écrit une matrice dans un fichier à partir d'une variable matrice déjà créée
func EcritureMat(taille int, mat [taille][taille]int, filename string) {
	file, err := os.Create(filename)

	if err != nil {
		log.Fatalln("Erreur lors de la création du fichier")
	}

	defer file.Close()

	for i := 0; i < taille; i++ {
		for j := 0; j < taille; j++ {
			_, err = fmt.Fprintf(file, "%d ", mat[i][j])
			if err != nil {
				log.Fatalln("Erreur lors de l'écriture dans le fichier")
			}
		}
		_, err = fmt.Fprintf(file, "\n")
		if err != nil {
			log.Fatalln("Erreur lors de l'écriture dans le fichier")
		}
	}
}
