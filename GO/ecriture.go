package main

import (
	"fmt"
	"log"
	"os"
)

// écrit une matrice dans un fichier à partir d'une variable matrice déjà créée
func EcritureMat(taille int, mat [taille]string, filename string) {
	file, err := os.Create(filename)

	if err != nil {
		log.Fatalln("Erreur lors de la création du fichier")
	}

	defer file.Close()

	for i := 0; i < taille; i++ {
		_, err = fmt.Fprintf(file, "%s ", mat[i])
		if err != nil {
			log.Fatalln("Erreur lors de l'écriture dans le fichier")
		}

		_, err = fmt.Fprintf(file, "\n")
		if err != nil {
			log.Fatalln("Erreur lors de l'écriture dans le fichier")
		}
	}
}
