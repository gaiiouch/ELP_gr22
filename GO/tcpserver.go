// go run tcpserver.go main.go ecriture.go lecture.go prodMat.go

package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr()
	fmt.Printf("Connection with %s established\n", remoteAddr)

	// creation des variables et des matrices
	var taille int
	var matA [][]int
	var matB [][]int
	var matRes [][]int

	// décodage des données reçues du client (bits -> matrices)
	decoder := gob.NewDecoder(conn)

	err := decoder.Decode(&taille) // taille des matrices
	if err != nil {
		fmt.Println("Erreur de réception :", err)
	}

	err = decoder.Decode(&matA) // première matrice
	if err != nil {
		fmt.Println("Erreur de réception :", err)
	}
	err = decoder.Decode(&matB) // deuxième matrice
	if err != nil {
		fmt.Println("Erreur de réception :", err)
	}

	//fmt.Println(matA)
	//fmt.Println(matB)

	// création des fichiers contenant les matrices
	nomFichier_1 := "matriceA.txt"
	nomFichier_2 := "matriceB.txt"
	nomFichier_3 := "matriceRes.txt"

	err = EcrireMatriceIntDansFichier(taille, matA, nomFichier_1)
	if err != nil {
		log.Fatalf("Erreur lors de l'écriture dans le fichier : %v", err)
	}

	err = EcrireMatriceIntDansFichier(taille, matB, nomFichier_2)
	if err != nil {
		log.Fatalf("Erreur lors de l'écriture dans le fichier : %v", err)
	}

	// calcul du produit des matrices à l'aide des fichiers
	err = Main(taille, nomFichier_1, nomFichier_2, nomFichier_3)
	if err != nil {
		fmt.Println("Erreur de calcul de la matrice résultat :", err)
	}

	// lecture de la matrice résultat dans le fichier correspondant
	matRes, err = LireMatriceDuFichier(nomFichier_3)
	if err != nil {
		fmt.Println("Erreur de lecture de la matrice :", err)
	}

	// encodage et envoi de la matrice résultat au client
	encoder := gob.NewEncoder(conn)

	err = encoder.Encode(matRes)
	if err != nil {
		log.Fatalf("Erreur lors de l'envoi de la matrice résultat : %v", err)
	}

	fmt.Printf("Connection with %s closed\n", remoteAddr)
}

func main() {

	ln, err := net.Listen("tcp", ":8000") // écoute sur le port 8000

	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handle(conn)
	}
}
