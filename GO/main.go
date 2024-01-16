// pour exécuter le programme : go run main.go lecture.go prodMat.go

// commande à lancer pour les tests : go run main.go lecture.go ecriture.go prodMat.go

package main

import (
	"log"
	"strconv"
	"sync"
)

// variables à ajouter : taille, matA, matB, matC, ligne, noms des fichiers
func Main(taille int, matA [taille][taille]int, matB [taille][taille]int, matC [taille][taille]int, ligne [taille]int, file_name []string) [taille][taille]int {

	// ouverture du wait group pour les go routines
	var wg sync.WaitGroup

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
		//fmt.Println("les data du channel", data)
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
			log.Fatalln("Erreur lors de la conversion en entier ligne 50")
		}

		// insertion du contenu de la ligne dans la matrice résultat
		c := k + 2
		d := 0
		for i := k + 3; i < len(data); i++ {
			if string(data[i]) == " " || string(data[i]) == "]" {
				//fmt.Println("variable c", c)
				//fmt.Println("variable i", i)

				val, err := strconv.Atoi(string(data[c:i])) // conversion de la première partie de la string en int (numéroDeLaLigne)
				if err != nil {
					log.Fatalln("Erreur lors de la conversion en entier ligne 60")
				}
				c = i + 1
				//fmt.Println("valeur res", val)
				ligne[d] = val
				d++
			}
		}
		matC[num_ligne] = ligne

	}

	wg.Wait()
	close(channel)

	// écriture du résultat dans une matrice pour ensuite l'envoyer au client
	err := EcritureMatInt(taille, matC, file_name[2])
	if err != nil {
		log.Fatalf("Erreur lors de l'écriture dans le fichier : %v", err)
	}
	//fmt.Println(matC)
	return matC
}

/*
func main() {

	var matA [taille][taille]int
	var matB [taille][taille]int
	var matRes [taille][taille]int

	for i := 0; i < taille; i++ {
		for j := 0; j < taille; j++ {
			matA[i][j] = rand.Intn(10)
			matB[i][j] = rand.Intn(10)
		}
	}
	//fmt.Println(matA)
	//fmt.Println(matB)
	var ligne [taille]int
	file_name := []string{"matriceA.txt", "matriceB.txt", "matriceResAB.txt"}
	matRes = Main(taille, matA, matB, matRes, ligne, file_name)
}
*/
