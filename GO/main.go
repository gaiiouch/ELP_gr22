// pour exécuter le programme : go run .

package main

import (
	"fmt"
)

func main() {

	matA, taille := LectureMat("matriceA.txt")
	fmt.Println(matA)

	var matB [taille][taille]int
	var matC [taille][taille]int

	for i := 0; i < taille; i++ {
		for j := 0; j < taille; j++ {
			matB[i][j] = 1
			matC[i][j] = 1
		}
	}

	matC = ProdMat(taille, matA, matB)
	fmt.Println(matC)

	err := EcritureMat(taille, matC, "matriceRes.txt")
	if err == 1 {
		fmt.Println("Erreur lors de l'écriture")
	}
}
