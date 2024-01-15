// go run tcpclient.go

package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"net"
)

func main() {

	// demande à l'utilisateur pour la taille des matrices
	var taille int
	fmt.Println("Enter a size for matrix :")
	fmt.Scan(&taille)

	// création des matrices avec des valeurs aléatoire
	var matA [][]int
	var matB [][]int
	var matRes [][]int

	matA = make([][]int, taille)
	matB = make([][]int, taille)
	for i := 0; i < taille; i++ {
		matA[i] = make([]int, taille)
		matB[i] = make([]int, taille)
		for j := 0; j < taille; j++ {
			matA[i][j] = rand.Intn(10)
			matB[i][j] = rand.Intn(10)
		}
	}

	// demande de connection au serveur
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// encodage des matrices pour les envoyer (matrice -> bits)
	encoder := gob.NewEncoder(conn)

	err = encoder.Encode(taille)
	if err != nil {
		log.Fatalf("Erreur lors de l'envoi de la taille : %v", err)
	}

	err = encoder.Encode(matA)
	if err != nil {
		log.Fatalf("Erreur lors de l'envoi de la matrice : %v", err)
	}

	err = encoder.Encode(matB)
	if err != nil {
		log.Fatalf("Erreur lors de l'envoi de la matrice : %v", err)
	}

	// décodage de la matrice envoyée par le serveur
	decoder := gob.NewDecoder(conn)

	err = decoder.Decode(&matRes)
	if err != nil {
		fmt.Println("Erreur de réception :", err)
	}

	fmt.Println(matRes)
}
