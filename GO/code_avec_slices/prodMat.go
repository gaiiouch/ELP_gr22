package main

import (
	"fmt"
	"sync"
)

// calcule le produit d'une ligne avec une matrice carrée
func ProduitMatrices(A [][]int, B [][]int, i int, channel chan string, wg *sync.WaitGroup) {

	defer wg.Done()
	n := len(A)
	resultat := make([]int, n)

	for j := 0; j < n; j++ {
		resultat[j] = 0
		for k := 0; k < n; k++ {
			resultat[j] = resultat[j] + A[i][k]*B[k][j]
		}
	}

	chaine := fmt.Sprintf("%d %d", i, resultat) // i = numéro de la ligne calculée ; resultat = contenu de la ligne
	channel <- chaine

}
