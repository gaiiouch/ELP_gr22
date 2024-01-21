package main

import (
	"sync"
)

type LigneMat struct {
	num_ligne     int
	contenu_ligne []int
}

func Main(taille int, matA [][]int, matB [][]int) ([][]int, error) {
	/*
		Exécute les goroutines pour la multiplication de 2 matrices
		Traite les résultats stockés dans un channel
		Retourne la matrice résultat de la multiplication
	*/

	var wg sync.WaitGroup

	nb_goroutines := taille
	channel := make(chan LigneMat)
	wg.Add(nb_goroutines)

	for i := 0; i < taille; i++ {
		go ProdMat(matA, matB, i, channel, &wg)
	}

	matC := make([][]int, taille)

	// pour chacune récupération de lignes dans le channel, on idenfie la ligne correspondante et on l'inclue dans la matrice résultat à la bonne position
	for j := 0; j < taille; j++ {
		data := <-channel
		matC[data.num_ligne] = data.contenu_ligne
	}

	wg.Wait()
	close(channel)

	return matC, nil
}
