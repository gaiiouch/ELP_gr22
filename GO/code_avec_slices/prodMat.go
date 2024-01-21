package main

import (
	"sync"
)

func ProdMat(A [][]int, B [][]int, i int, channel chan LigneMat, wg *sync.WaitGroup) {
	/*
	 calcule le produit d'une ligne avec une matrice carrée
	*/

	defer wg.Done()
	n := len(A)
	C := make([]int, n)

	for j := 0; j < n; j++ {
		C[j] = 0
		for k := 0; k < n; k++ {
			C[j] = C[j] + A[i][k]*B[k][j]
		}
	}
	l := LigneMat{num_ligne: i, contenu_ligne: C}
	//chaine := fmt.Sprintf("%d %d", i, resultat) // i = numéro de la ligne calculée ; resultat = contenu de la ligne
	channel <- l

}
