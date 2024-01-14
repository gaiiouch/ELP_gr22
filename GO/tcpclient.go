package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer conn.Close()

	for i := 0; i < 5; i++ {
		io.WriteString(conn, fmt.Sprintf("Coucou %d\n", i))
	}

	conn.Close()
}
