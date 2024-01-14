package main

import (
	"fmt"
	"sync"
)

func ProduitMatrices(A [][]int, B [][]int, a int, b int, channel chan string, wg *sync.WaitGroup) {

	defer wg.Done()
	// produit de matrices carrées donc le résultat à la même taille que les matrices servant à effectuer le calcul
	n := len(A)
	resultat := make([]int, n)

	for j := 0; j < n; j++ {
		resultat[j] = 0

		for k := 0; k < n; k++ {
			resultat[j] = resultat[j] + A[a][k]*B[k][j]
		}

	}

	chaine := fmt.Sprintf("%d %d", a, resultat)
	channel <- chaine

}
