// FICHIER INUTILE

package main

import (
	"bufio"
	"os"
	"strconv"
)

// lit une matrice dans un fichier pour la transformer en variable pour le code
func LectureMat(taille int, mat [taille][taille]int, filename string) ([taille][taille]int, error) {

	file, err := os.Open(filename)

	if err != nil {
		return mat, err
	}

	defer file.Close() // assurer la fermeture du fichier si le programme plante

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords) // lecture du fichier mot par mot (ici : nombre par nombre)

	for i := 0; i < taille; i++ {

		for j := 0; j < taille; j++ {

			scanner.Scan()
			num, err := strconv.Atoi(scanner.Text()) // conversion de string en int

			if err != nil {
				return mat, err
			}

			mat[i][j] = num // ajout à la bonne position dans la matrice
		}
	}
	return mat, nil
}
