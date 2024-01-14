// go run tcpserver.go

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

	// creation des matrices
	var matA [taille][taille]int
	var matB [taille][taille]int

	// décoder les données reçues du client (bites -> matrices)
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

	fmt.Printf("Connection with %s closed\n", remoteAddr)

	/*scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		data := scanner.Text()
		fmt.Println(data)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading:", err)
	}*/
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
