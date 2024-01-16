// FICHIER INUTILE

package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func LectureMat(filename string) ([][]int, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// calcul de la taille de la matrice pour pouvoir la créer ensuite
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var nb_valeurs float64 = 0

	for scanner.Scan() {
		nb_valeurs += 1
	}

	var n int = int(math.Sqrt(nb_valeurs))

	// réinitialisation du scanner pour pouvoir remplir la matrice en déplacant le curseur au début du fichier
	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	// remplissage de la matrice
	mat := make([][]int, n)

	for i := 0; i < n; i++ {

		mat[i] = make([]int, n)

		for j := 0; j < n; j++ {

			scanner.Scan()
			num, err := strconv.Atoi(scanner.Text())

			if err != nil {
				return nil, err
			}

			mat[i][j] = num
		}
	}

	return mat, nil
}
