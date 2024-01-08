package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

// lit une matrice dans un fichier pour la transformer en variable pour le code
func LectureMat(filename string) ([taille][taille]int, const) {

	file, err := os.Open(filename)

	if err != nil {
		os.Exit(0)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var nb_val float64 = 0

	for scanner.Scan() {
		nb_val += 1
	}

	const taille int = int(math.Sqrt(nb_val))

	var mat [taille][taille]int

	for i := 0; i < taille; i++ {

		for j := 0; j < taille; j++ {

			num, err := strconv.Atoi(scanner.Text())

			if err != nil {
				os.Exit(0)
			}

			mat[i][j] = num
		}
	}
	return mat, taille
}
