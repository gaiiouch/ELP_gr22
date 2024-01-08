package main

import (
	"fmt"
	"log"
)

func main() {
	// lecture de la première matrice dans un fichier
	nomFichier_1 := "matriceA.txt"

	matriceA, err := LireMatriceDuFichier(nomFichier_1)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier : %v", err)
	}

	fmt.Println("Matrice lue depuis le fichier :", matriceA)

	// création de la deuxième matrice
	matriceB := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	// produit des deux matrices avec vérification que les matrices soient carrées et de même taille
	if len(matriceA) != len(matriceB) || len(matriceA) == 0 {
		fmt.Println("Les matrices ne sont pas de taille carrée ou sont vides.")
		return
	}

	produit := ProduitMatrices(matriceA, matriceB)
	fmt.Println("Produit des matrices :", produit)

	// ecriture du résultat du produit dans un fichier
	nomFichier_2 := "matriceRes.txt"

	err = EcrireMatriceDansFichier(len(produit), produit, nomFichier_2)
	if err != nil {
		log.Fatalf("Erreur lors de l'écriture dans le fichier : %v", err)
	}

	fmt.Printf("Matrice écrite dans le fichier %s\n", nomFichier_2)
}
