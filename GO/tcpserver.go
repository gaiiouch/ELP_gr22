package main

import (
	"bufio"
	"fmt"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()

	for {
		data, err := bufio.NewReader(conn).ReadString('\n') //read the data only string
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(data)
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
