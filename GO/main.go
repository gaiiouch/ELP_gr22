package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)

func main() {

	// ouverture du wait group pour les go routines
	var wg sync.WaitGroup

	// lecture des matrices dans des fichiers
	nomFichier_1 := "matriceA.txt"
	matriceA, err := LireMatriceDuFichier(nomFichier_1)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier %s : %v", nomFichier_1, err)
	}

	nomFichier_2 := "matriceB.txt"
	matriceB, err := LireMatriceDuFichier(nomFichier_2)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier %s : %v", nomFichier_2, err)
	}

	// produit des deux matrices avec vérification que les matrices soient carrées et de même taille
	if len(matriceA) != len(matriceB) || len(matriceA) == 0 {
		log.Fatalf("Les matrices ne sont pas de taille carrée ou sont vides.")
	}
	taille := len(matriceA)

	// préparation pour les goroutines avec canal pour stocker chaque ligne de la matrice calculée
	a := 0
	b := 1
	nb_goroutines := taille
	channel := make(chan string)
	wg.Add(nb_goroutines) // nb de goroutines à attendre

	// pour chaque ligne de la première matrice, on calcule via les goroutines la ligne correspondante pour la matrice résultat
	for i := 0; i < taille; i++ {
		go ProduitMatrices(matriceA, matriceB, a, b, channel, &wg)
		a++
		b++
	}

	// création du tableau pour le résultat final
	produit := make([]string, taille)

	// pour chacune récupération de lignes dans le channel, on idenfie la ligne correspondante et on l'inclue dans la matrice résultat à la bonne position
	for v := 0; v < taille; v++ {
		u := <-channel // format de u : "numéroDeLaLigne [contenuDeLaLigne]"

		k := 0
		for {
			if string(u[k]) == " " { // recherche du premier espace dans des strings dans le canal
				break
			}
			k++
		}

		num_ligne, err := strconv.Atoi(string(u[:k])) // conversion de la première partie de la string en int (numéroDeLaLigne)
		if err != nil {
			log.Fatalln("Erreur lors de la conversion en entier")
		}

		fmt.Println(num_ligne)

		// insertion du contenu de la ligne dans la matrice résultat
		ligne := string(u[k+2 : len(u)-1])
		produit[num_ligne] = ligne

	}

	wg.Wait()
	close(channel)

	// ecriture du résultat du produit dans un fichier
	nomFichier_3 := "matriceRes.txt"
	err = EcrireMatriceDansFichier(len(produit), produit, nomFichier_3)
	if err != nil {
		log.Fatalf("Erreur lors de l'écriture dans le fichier : %v", err)
	}
}
