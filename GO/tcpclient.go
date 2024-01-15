// go run tcpclient.go

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

	// connection avec le serveur
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// encodage des données à envoyer au serveur (matrices -> bits)
	encoder := gob.NewEncoder(conn)
	err = encoder.Encode(matA)
	if err != nil {
		log.Fatalf("Erreur lors de l'envoi de la matrice : %v", err)
	}

	err = encoder.Encode(matB)
	if err != nil {
		log.Fatalf("Erreur lors de l'envoi de la matrice : %v", err)
	}

	// décodage des données envoyées par le serveur
	decoder := gob.NewDecoder(conn)

	err = decoder.Decode(&matRes) // matrice résultat
	if err != nil {
		fmt.Println("Erreur de réception :", err)
	}

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
	//fmt.Println(matRes)

}
