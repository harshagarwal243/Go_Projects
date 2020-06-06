package main

import (
	"fmt"
	"net"
	"bufio"
	"log"
	"os"
)

func main() {
	li , err := net.Listen("tcp",":3031")
	if err != nil {
		log.Panic(err)
	}
	for {
		acc , err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handle(acc)
	}
}

func handle(con net.Conn) {
	defer con.Close()
	scanner := bufio.NewScanner(con)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println("Message from Client : "+ln)
		message , _ := bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Fprintf(con,"This is from Server Side : "+message)
	}
}