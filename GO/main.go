// pour exécuter le programme : go run .

package main

import (
	"fmt"
)

const taille int = 3

func main() {

	var matA [taille][taille]int
	var matB [taille][taille]int
	var matC [taille][taille]int

	matA = LectureMat(taille, matA, "matriceA.txt")
	fmt.Println(matA)

	for i := 0; i < taille; i++ {
		for j := 0; j < taille; j++ {
			matB[i][j] = 1
			matC[i][j] = 1
		}
	}

	matC = ProdMat(taille, matA, matB)
	fmt.Println(matC)

	EcritureMat(taille, matC, "matriceRes.txt")

}
