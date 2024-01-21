// go run tcpserver.go main.go prodMat.go

package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func handle(conn net.Conn) {
	/*
		Réceptionne les matrices envoyées par le client et lance le calcul du résultat du produit.
		Envoie le résultat au client.
	*/

	defer conn.Close()

	remoteAddr := conn.RemoteAddr()
	fmt.Printf("Connection with %s established\n", remoteAddr)

	var taille int
	var matA, matB, matRes [][]int

	decoder := gob.NewDecoder(conn)
	err := decoder.Decode(&taille)
	if err != nil {
		fmt.Println("	Erreur de réception :", err)
		return
	}
	err = decoder.Decode(&matA)
	if err != nil {
		fmt.Println("	Erreur de réception :", err)
		return
	}
	err = decoder.Decode(&matB)
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

	fmt.Printf("Connection with %s closed\n", remoteAddr)
}

func main() {
	/*
		Met en place un serveur TCP qui attend d'établir le lien avec un client.
	*/

	ln, err := net.Listen("tcp", ":8000")
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
