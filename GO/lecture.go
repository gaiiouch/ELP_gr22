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
	/*
		scanner2 := bufio.NewScanner((file))
		scanner2.Split(bufio.ScanLines)

		var nb_val int = 0

		for scanner2.Scan() {
			nb_val += 1
		}

		const taille int = nb_val

		var mat [taille][taille]int
	*/

	scanner.Scan()

	for i := 0; i < taille; i++ {

		for j := 0; j < taille; j++ {

			num, err := strconv.Atoi(scanner.Text())

			if err != nil {
				log.Fatalln("Erreur lors de la conversion en entier")
			}

			mat[i][j] = num
		}
	}
	return mat
}
