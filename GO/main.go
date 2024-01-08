package main

import (
	"fmt"
)

const taille int = 3

func main() {

	var matA [taille][taille]int
	var matB [taille][taille]int
	var matC [taille][taille]int

	matA, err := LectureMat(taille, matA, "matriceA.txt")
	if err == 1 {
		fmt.Println("Erreur lors de la lecture")
	}
	fmt.Println(matA)

	for i := 0; i < taille; i++ {
		for j := 0; j < taille; j++ {
			matB[i][j] = 1
			matC[i][j] = 1
		}
	}

	matC = ProdMat(taille, matA, matB)
	fmt.Println(matC)

	err = EcritureMat(taille, matC, "matriceRes.txt")
	if err == 1 {
		fmt.Println("Erreur lors de l'écriture")
	}
}
