// pour exécuter le programme : go run main.go lecture.go prodMat.go

package main

import (
	"log"
	"strconv"
	"sync"
)

// variables à ajouter : taille, matA, matB, matC, ligne, noms des fichiers
func Main(taille int, matA [taille][taille]int, matB [taille][taille]int, matC [taille]string, ligne [taille]int, file_name []string) {

	// ouverture du wait group pour les go routines
	var wg sync.WaitGroup

	// lecture des fichiers contenant les matrices
	matA, err := LectureMat(taille, matA, file_name[0])
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier %s : %v", file_name, err)
	}

	matB, err = LectureMat(taille, matB, file_name[1])
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier %s : %v", file_name, err)
	}

	// préparation pour les goroutines avec canal pour stocker chaque ligne de la matrice calculée
	a := 0
	b := 1 // ACTUELLEMENT INUTILE SI ON FAIT LIGNE PAR LIGNE
	nb_goroutines := taille
	channel := make(chan string)
	wg.Add(nb_goroutines) // nb de goroutines à attendre

	// pour chaque ligne de la première matrice, on calcule via les goroutines la ligne correspondante pour la matrice résultat
	for i := 0; i < taille; i++ {
		go ProdMat(taille, matA, matB, ligne, a, b, channel, &wg)
		a++
		b++
	}

	// pour chacune récupération de lignes dans le channel, on idenfie la ligne correspondante et on l'inclue dans la matrice résultat à la bonne position
	for j := 0; j < taille; j++ {
		data := <-channel // format de data : "numéroDeLaLigne [contenuDeLaLigne]"

		k := 0
		// recherche du premier espace dans les strings du canal
		for {
			if string(data[k]) == " " {
				break
			}
			k++
		}

		num_ligne, err := strconv.Atoi(string(data[:k])) // conversion de la première partie de la string en int (numéroDeLaLigne)
		if err != nil {
			log.Fatalln("Erreur lors de la conversion en entier")
		}

		// insertion du contenu de la ligne dans la matrice résultat
		ligne := string(data[k+2 : len(data)-1])
		matC[num_ligne] = ligne

	}

	wg.Wait()
	close(channel)

	// écriture du résultat dans une matrice pour ensuite l'envoyer au client
	err = EcritureMatString(taille, matC, file_name[2])
	if err != nil {
		log.Fatalf("Erreur lors de l'écriture dans le fichier : %v", err)
	}

}
