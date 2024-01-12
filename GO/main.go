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
	nb_goroutines := 3
	channel := make(chan string, 2)
	wg.Add(nb_goroutines)
	for i := 0; i < taille; i++ {
		go ProdMat(taille, matA, matB, matC, a, b, channel, &wg)
		a++
		b++
		v := <-channel
		fmt.Println(v)
	}

	wg.Wait()
	close(channel)

	//matC = ProdMat(taille, matA, matB, a, b)
	//fmt.Println(matC)

	EcritureMat(taille, matC, "matriceRes.txt")

}
