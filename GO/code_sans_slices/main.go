package main

import (
	"fmt"
	"strconv"
	"sync"
)

// exécute les goroutines pour la multiplication de 2 matrices et traitement des résultats stockés dans un channel
func Main(taille int, matA [taille][taille]int, matB [taille][taille]int, matC [taille][taille]int, ligne [taille]int) ([taille][taille]int, error) {

	// ouverture du wait group pour les go routines
	var wg sync.WaitGroup

	// préparation pour les goroutines avec canal pour stocker chaque ligne de la matrice calculée
	nb_goroutines := taille
	channel := make(chan string)
	wg.Add(nb_goroutines) // nb de goroutines à attendre

	// pour chaque ligne de la première matrice, on calcule via les goroutines la ligne correspondante pour la matrice résultat
	for i := 0; i < taille; i++ {
		go ProdMat(taille, matA, matB, ligne, i, channel, &wg)
	}

	// pour chacune des lignes récupérées dans le channel, on idenfie la ligne correspondante et on l'inclue dans la matrice résultat à la bonne position
	for j := 0; j < taille; j++ {
		data := <-channel // format de data : "numéroDeLaLigne [contenuDeLaLigne]"

		// recherche du premier espace dans les strings du canal
		k := 0
		for {
			if string(data[k]) == " " {
				break
			}
			k++
		}

		num_ligne, err := strconv.Atoi(string(data[:k])) // conversion de la première partie de la string en int (numéroDeLaLigne)
		if err != nil {
			fmt.Println("Erreur lors de la conversion en entier (ligne 38) :", err)
			return matC, err
		}

		// insertion du contenu de la ligne dans la matrice résultat
		x := k + 2 // premier endroit logique où on trouve un début de nombre
		y := 0
		for i := k + 3; i < len(data); i++ {
			if string(data[i]) == " " || string(data[i]) == "]" {
				val, err := strconv.Atoi(string(data[x:i])) // conversion de la valeur trouvée en int
				if err != nil {
					fmt.Println("Erreur lors de la conversion en entier (ligne 50) :", err)
					return matC, err
				}
				x = i + 1 // prochain endroit logique où on trouve un début de nombre
				ligne[y] = val
				y++
			}
		}
		matC[num_ligne] = ligne // ajout de la ligne à la matrice résultat
	}

	wg.Wait()      // attendre que toutes les goroutines soient finies
	close(channel) //fermeture du channel

	return matC, nil
}
