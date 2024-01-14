// go run tcpserver.go

package main

import (
	"bufio"
	"fmt"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr()
	fmt.Printf("Connection established with %s\n", remoteAddr)

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		data := scanner.Text()
		fmt.Println(data)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading:", err)
	}
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
