package main

import (
	"fmt"
	"strconv"
	"sync"
)

func Main(taille int, matA [taille][taille]int, matB [taille][taille]int, matC [taille][taille]int, ligne [taille]int) ([taille][taille]int, error) {
	/*
		Exécute les goroutines pour la multiplication de 2 matrices
		Traite les résultats stockés dans un channel
		Retourne la matrice résultat de la multiplication
	*/

	type LigneMat struct {
		num_ligne int
		ligne     [taille]int
	}

	var wg sync.WaitGroup

	nb_goroutines := taille
	channel := make(chan string)
	wg.Add(nb_goroutines)

	for i := 0; i < taille; i++ {
		go ProdMat(taille, matA, matB, ligne, i, channel, &wg)
	}

	// pour chacune des lignes récupérées dans le channel, on idenfie la ligne correspondante et on l'inclue dans la matrice résultat à la bonne position
	for j := 0; j < taille; j++ {
		data := <-channel

		k := 0
		for {
			if string(data[k]) == " " {
				break
			}
			k++
		}

		num_ligne, err := strconv.Atoi(string(data[:k]))
		if err != nil {
			fmt.Println("Erreur lors de la conversion en entier (ligne 38) :", err)
			return matC, err
		}

		// convertion et insertion du contenu de la ligne dans la matrice résultat
		x := k + 2
		y := 0
		for i := k + 3; i < len(data); i++ {
			if string(data[i]) == " " || string(data[i]) == "]" {
				val, err := strconv.Atoi(string(data[x:i]))
				if err != nil {
					fmt.Println("Erreur lors de la conversion en entier (ligne 50) :", err)
					return matC, err
				}
				x = i + 1
				ligne[y] = val
				y++
			}
		}
		matC[num_ligne] = ligne
	}

	wg.Wait()
	close(channel)

	return matC, nil
}
