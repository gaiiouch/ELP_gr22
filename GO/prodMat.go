package main

import "fmt"

const taille int = 3

// calcule le produit de deux matrices carrées de même taille
func prodMat(taille int, A [taille][taille]int, B [taille][taille]int) (C [taille][taille]int) {

	for i := 0; i < taille; i++ {
		for j := 0; j < taille; j++ {
			C[i][j] = 0
			for k := 0; k < taille; k++ {
				C[i][j] = C[i][j] + A[i][k]*B[k][j]
			}
		}
	}
	return
}

func main() {

	var matA [taille][taille]int
	var matB [taille][taille]int
	var matC [taille][taille]int

	for i := 0; i < taille; i++ {
		for j := 0; j < taille; j++ {
			matA[i][j] = 2
			matB[i][j] = 1
			matC[i][j] = 1
		}
	}

	matC = prodMat(taille, matA, matB)
	fmt.Println(matC)
}
