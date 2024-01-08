package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// lit une matrice dans un fichier pour la transformer en variable pour le code
func LectureMat(taille int, mat [taille][taille]int, filename string) [taille][taille]int {

	file, err := os.Open(filename)

	if err != nil {
		log.Fatalln("Erreur lors de l'ouverture du fichier")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for i := 0; i < taille; i++ {

		for j := 0; j < taille; j++ {

			scanner.Scan()
			num, err := strconv.Atoi(scanner.Text())

			if err != nil {
				log.Fatalln("Erreur lors de la conversion en entier")
			}

			mat[i][j] = num
		}
	}
	return mat
}
