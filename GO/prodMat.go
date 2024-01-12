package main

import (
	"fmt"
	"sync"
)

// calcule le produit de deux matrices carrées de même taille
func ProdMat(taille int, A [taille][taille]int, B [taille][taille]int, C [taille][taille]int, a int, b int, channel chan string, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := a; i < b; i++ {
		for j := 0; j < taille; j++ {
			C[i][j] = 0
			for k := 0; k < taille; k++ {
				C[i][j] = C[i][j] + A[i][k]*B[k][j]
			}
		}
	}
	chaine := fmt.Sprintf("%d %d", a, C)
	channel <- chaine

}
