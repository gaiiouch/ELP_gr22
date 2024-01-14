// go run tcpclient.go ecriture.go

package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"net"
)

const taille int = 9

func main() {

	var matA [taille][taille]int
	var matB [taille][taille]int

	for i := 0; i < taille; i++ {
		for j := 0; j < taille; j++ {
			matA[i][j] = rand.Intn(10)
			matB[i][j] = rand.Intn(10)
		}
	}

	//fmt.Println(matA)
	//fmt.Println(matB)

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	encoder := gob.NewEncoder(conn)
	err = encoder.Encode(matA)
	if err != nil {
		log.Fatalf("Erreur lors de l'envoi de la matrice : %v", err)
	}

	err = encoder.Encode(matB)
	if err != nil {
		log.Fatalf("Erreur lors de l'envoi de la matrice : %v", err)
	}

	/*
		nom_fichier1 := "matriceA.txt"
		err = EcritureMatInt(taille, matA, nom_fichier1)
		if err != nil {
			log.Fatalf("Erreur lors de l'écriture dans le fichier : %v", err)
		}

		nom_fichier2 := "matriceB.txt"
		err = EcritureMatInt(taille, matB, nom_fichier2)
		if err != nil {
			log.Fatalf("Erreur lors de l'écriture dans le fichier : %v", err)
		}

		io.WriteString(conn, fmt.Sprintf("Matrices écrites dans les fichiers %s et %s", nom_fichier1, nom_fichier2))*/

}
