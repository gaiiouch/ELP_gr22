package main

import (
	"bufio"
	"fmt"
	"net"
)

func main(){
	//listen 
	//accept 
	//handle connection with threads 

	dstream ,err := net.Listen("tcp", ":8000")

	if err != nil {
		fmt.Println(err)
		return 
	}

	defer dstream.Close()

	for{
		con,err := dstream.Accept()
		if err != nil {
			fmt.Println(err)
			return 
		}
		go handle(con)
	}
}

func handle(con net.Conn) {
	for {
		data ,err := bufio.NewReader(con).ReadString('\n') //read the data only string 
		if err != nil {
			fmt.Println(err)
			return 
		}
		fmt.Println(data)
	}
	con.Close()
}