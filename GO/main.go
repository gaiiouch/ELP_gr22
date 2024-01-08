package main

import (
	"fmt"
)

func main() {
	const taille = 3

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

	matC = ProdMat(taille, matA, matB)
	fmt.Println(matC)

}
