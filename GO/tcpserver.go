// go run tcpserver.go main.go lecture.go ecriture.go prodMat.go

package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

const taille int = 9

func handle(conn net.Conn) {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr()
	fmt.Printf("Connection established with %s\n", remoteAddr)

	// creation des variables et des matrices
	var matA [taille][taille]int
	var matB [taille][taille]int
	var matRes [taille][taille]int
	var matC [taille]string
	var ligne [taille]int
	file_name := []string{"matriceA.txt", "matriceB.txt", "matriceRes.txt"}

	// décoder les données reçues du client (bits -> matrices)
	decoder := gob.NewDecoder(conn)

	err := decoder.Decode(&matA) // Première matrice
	if err != nil {
		fmt.Println("Erreur de réception :", err)
	}
	err = decoder.Decode(&matB) // Deuxième matrice
	if err != nil {
		fmt.Println("Erreur de réception :", err)
	}

	//fmt.Println(matA)
	//fmt.Println(matB)
	err = EcritureMatInt(taille, matA, file_name[0])
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier : %v", err)
	}

	err = EcritureMatInt(taille, matB, file_name[1])
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier : %v", err)
	}

	// calcul du produit
	Main(taille, matA, matB, matC, ligne, file_name)

	// lecture du fichier où se trouve le résultat
	matRes, err = LectureMat(taille, matRes, file_name[2])
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier %s : %v", file_name, err)
	}

	// encodage du résultat pour l'envoi au client
	encoder := gob.NewEncoder(conn)
	err = encoder.Encode(matRes)
	if err != nil {
		fmt.Printf("Erreur lors de l'envoi de la matrice : %v", err)
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
