// go run tcpclient.go ecriture.go

package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

const taille int = 15

func main() {

	start := time.Now()

	// creation des variables et des matrices (générées aléatoirement)
	var matA, matB, matRes [taille][taille]int
	file_name := []string{"matriceA.txt", "matriceB.txt", "matriceResAB.txt"}

	for i := 0; i < taille; i++ {
		for j := 0; j < taille; j++ {
			matA[i][j] = rand.Intn(10)
			matB[i][j] = rand.Intn(10)
		}
	}

	// connexion avec le serveur
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalln("Erreur lors de la connexion avec le serveur :", err)
	}
	defer conn.Close()

	// encodage des données à envoyer au serveur (matrices -> bits)
	encoder := gob.NewEncoder(conn)
	err = encoder.Encode(matA) // première matrice
	if err != nil {
		log.Fatalln("Erreur lors de l'envoi de la matrice :", err)
	}
	err = encoder.Encode(matB) // deuxième matrice
	if err != nil {
		log.Fatalln("Erreur lors de l'envoi de la matrice :", err)
	}

	// décodage des données envoyées par le serveur
	decoder := gob.NewDecoder(conn)
	err = decoder.Decode(&matRes) // matrice résultat
	if err != nil {
		log.Fatalln("Erreur de réception :", err)
	}

	// temps pour génération des matrices, envoi au serveur, calcul du produit, réception du serveur et décodage du résultat
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)

	// écriture de chacune des matrices dans un fichier texte (optionnelle)
	err = EcritureMat(taille, matA, file_name[0])
	if err != nil {
		log.Fatalln("Erreur lors de l'écriture dans le fichier :", err)
	}
	err = EcritureMat(taille, matB, file_name[1])
	if err != nil {
		log.Fatalln("Erreur lors de l'écriture dans le fichier :", err)
	}
	err = EcritureMat(taille, matRes, file_name[2])
	if err != nil {
		log.Fatalln("Erreur lors de l'écriture dans le fichier :", err)
	}
}
