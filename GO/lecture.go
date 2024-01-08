package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func LireMatriceDuFichier(nomFichier string) ([][]int, error) {

	fichier, err := os.Open(nomFichier)
	if err != nil {
		return nil, err
	}
	defer fichier.Close()

	// calcul de la taille de la matrice pour pouvoir la créer ensuite
	scanner := bufio.NewScanner(fichier)
	scanner.Split(bufio.ScanWords)

	var nb_valeurs float64 = 0

	for scanner.Scan() {
		nb_valeurs += 1
	}

	var n int = int(math.Sqrt(nb_valeurs))

	matrice := make([][]int, n)

	// réinitialisation du scanner pour pouvoir remplir la matrice en déplacant le curseur au début du fichier
	_, err = fichier.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	scanner = bufio.NewScanner(fichier)
	scanner.Split(bufio.ScanWords)

	// remplissage de la matrice
	for i := 0; i < n; i++ {

		matrice[i] = make([]int, n)

		for j := 0; j < n; j++ {

			scanner.Scan()
			val, err := strconv.Atoi(scanner.Text())

			if err != nil {
				return nil, err
			}

			matrice[i][j] = val
		}
	}

	return matrice, nil
}
