// pour exécuter le programme : go run .

package main

import (
	"fmt"
	"sync"
)

const taille int = 3

func main() {

	var wg sync.WaitGroup

	var matA [taille][taille]int
	var matB [taille][taille]int
	var matC [taille][taille]int

	matA = LectureMat(taille, matA, "matriceA.txt")
	fmt.Println(matA)

	for i := 0; i < taille; i++ {
		for j := 0; j < taille; j++ {
			matB[i][j] = 1
		}
	}

	a := 0
	b := 1

	channel := make(chan string)

	for i := 0; i < taille; i++ {
		wg.Add(1)
		go ProdMat(taille, matA, matB, matC, a, b, channel, &wg)
		a++
		b++
	}

	//close(channel)
	wg.Wait()

	for valeur := range channel {
		fmt.Println(valeur)
	}

	//matC = ProdMat(taille, matA, matB, a, b)
	//fmt.Println(matC)

	EcritureMat(taille, matC, "matriceRes.txt")

}
