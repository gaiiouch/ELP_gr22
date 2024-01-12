// pour exécuter le programme : go run .

package main

import (
	"log"
	"strconv"
	"sync"
)

const taille int = 9

func main() {

	var wg sync.WaitGroup

	var matA [taille][taille]int
	var matB [taille][taille]int

	matA = LectureMat(taille, matA, "matriceA.txt")
	matB = LectureMat(taille, matB, "matriceB.txt")

	a := 0
	b := 1
	nb_goroutines := taille
	channel := make(chan string)
	wg.Add(nb_goroutines)

	for i := 0; i < taille; i++ {
		var ligne [taille]int
		go ProdMat(taille, matA, matB, ligne, a, b, channel, &wg)
		a++
		b++
	}

	var matC [taille]string

	for v := 0; v < taille; v++ {
		u := <-channel

		k := 0
		for {
			if string(u[k]) == " " {
				break
			}
			k++
		}

		num_ligne, err := strconv.Atoi(string(u[:k]))

		if err != nil {
			log.Fatalln("Erreur lors de la conversion en entier")
		}

		//fmt.Println(u, num_ligne)

		ligne := string(u[k+2 : len(u)-1])
		//fmt.Println(ligne)

		matC[num_ligne] = ligne

	}

	wg.Wait()
	close(channel)

	//matC = ProdMat(taille, matA, matB, a, b)
	//fmt.Println(matC)

	EcritureMat(taille, matC, "matriceRes.txt")

}
