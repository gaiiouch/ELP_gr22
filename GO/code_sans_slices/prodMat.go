package main

import (
	"sync"
)

func ProdMat(taille int, A [taille][taille]int, B [taille][taille]int, C [taille]int, i int, channel chan LigneMat, wg *sync.WaitGroup) {
	/*
	 calcule le produit d'une ligne avec une matrice carrée
	*/

	defer wg.Done()

	for j := 0; j < taille; j++ {
		C[j] = 0
		for k := 0; k < taille; k++ {
			C[j] = C[j] + A[i][k]*B[k][j]
		}
	}
	l := LigneMat{num_ligne: i, contenu_ligne: C}
	// chaine := fmt.Sprintf("%d %d", i, C) // i = numéro de la ligne calculée ; C = contenu de la ligne
	channel <- l
}
