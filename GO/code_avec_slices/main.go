package main

import (
	"fmt"
	"strconv"
	"sync"
)

func Main(taille int, matA [][]int, matB [][]int) ([][]int, error) {

	// ouverture du wait group pour les go routines
	var wg sync.WaitGroup

	// préparation pour les goroutines avec canal pour stocker chaque ligne de la matrice calculée
	nb_goroutines := taille
	channel := make(chan string)
	wg.Add(nb_goroutines) // nb de goroutines à attendre

	// pour chaque ligne de la première matrice, on calcule via les goroutines la ligne correspondante pour la matrice résultat
	for i := 0; i < taille; i++ {
		go ProduitMatrices(matA, matB, i, channel, &wg)
	}

	// création du tableau pour le résultat final
	ligne := make([]int, taille)
	matC := make([][]int, taille)

	// pour chacune récupération de lignes dans le channel, on idenfie la ligne correspondante et on l'inclue dans la matrice résultat à la bonne position
	for j := 0; j < taille; j++ {
		data := <-channel // format de data : "numéroDeLaLigne [contenuDeLaLigne]"

		k := 0
		for {
			if string(data[k]) == " " { // recherche du premier espace dans des strings dans le canal
				break
			}
			k++
		}

		num_ligne, err := strconv.Atoi(string(data[:k])) // conversion de la première partie de la string en int (numéroDeLaLigne)
		if err != nil {
			fmt.Println("Erreur lors de la conversion en entier :", err)
			return nil, err
		}

		// insertion du contenu de la ligne dans la matrice résultat
		x := k + 2 // premier endroit logique où on trouve un début de nombre
		y := 0
		for i := k + 3; i < len(data); i++ {
			if string(data[i]) == " " || string(data[i]) == "]" {

				val, err := strconv.Atoi(string(data[x:i])) // conversion de la valeur trouvée en int
				if err != nil {
					fmt.Println("Erreur lors de la conversion en entier (ligne 49) :", err)
					return nil, err
				}
				x = i + 1 // prochain endroit logique où on trouve un début de nombre
				ligne[y] = val
				y++
			}
		}
		matC[num_ligne] = ligne // ajout de la ligne à la matrice résultat

	}

	wg.Wait()
	close(channel)

	return matC, nil
}
