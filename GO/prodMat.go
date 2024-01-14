package main

import (
	"fmt"
	"sync"
)

// calcule le produit de d'une ligne avec une matrice carrée
func ProdMat(taille int, A [taille][taille]int, B [taille][taille]int, C [taille]int, a int, b int, channel chan string, wg *sync.WaitGroup) {

	defer wg.Done()
	for j := 0; j < taille; j++ {
		C[j] = 0
		for k := 0; k < taille; k++ {
			C[j] = C[j] + A[a][k]*B[k][j]
		}
	}
	chaine := fmt.Sprintf("%d %d", a, C) // a = numéro de la ligne ; C = contenu de la ligne
	channel <- chaine
}
