// go run tcpserver.go main.go prodMat.go

package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()

	// annonce du début d'une connexion
	remoteAddr := conn.RemoteAddr()
	fmt.Printf("Connection with %s established\n", remoteAddr)

	// creation des variables et des matrices
	var taille int
	var matA, matB, matRes [][]int

	// décodage des données reçues du client (bits -> matrices)
	decoder := gob.NewDecoder(conn)
	err := decoder.Decode(&taille) // taille
	if err != nil {
		fmt.Println("	Erreur de réception :", err)
		return
	}
	err = decoder.Decode(&matA) // première matrice
	if err != nil {
		fmt.Println("	Erreur de réception :", err)
		return
	}
	err = decoder.Decode(&matB) // deuxième matrice
	if err != nil {
		fmt.Println("	Erreur de réception :", err)
		return
	}

	// calcul du produit des matrices
	matRes, err = Main(taille, matA, matB)
	if err != nil {
		fmt.Println("	Erreur de calcul de la matrice résultat :", err)
		return
	} else {
		fmt.Println("	Calcul du produit des matrices réussi")
	}

	// encodage du résultat pour l'envoi au client
	encoder := gob.NewEncoder(conn)
	err = encoder.Encode(matRes)
	if err != nil {
		fmt.Printf("	Erreur lors de l'envoi de la matrice : %v", err)
		return
	}

	// annonce de la fin de la connexion
	fmt.Printf("Connection with %s closed\n", remoteAddr)
}

func main() {
	// écoute sur le port 8000
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	// acceptation et gestion des connexions entrantes
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handle(conn)
	}
}