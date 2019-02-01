package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(c net.Conn, e net.Conn) {

	reader := bufio.NewReader(c)

	defer c.Close()
	fmt.Printf("Connected to : %v\n", c.RemoteAddr())

	for {
		text, _ := reader.ReadString('\n')
		// send
		fmt.Fprintf(e, text+"\n")
		// listen & return message
		m, _ := bufio.NewReader(e).ReadString('\n')
		fmt.Print("Message from server: " + m)
		// passes msg along
		fmt.Fprintf(c, "%s\r\n", m)
	}
}

func main() {

	fmt.Println("relay runs")

	var echoServerURL = "localhost:8080"

	port := "8081"
	fmt.Println("Listening on Port", port)

	e, _ := net.Dial("tcp", echoServerURL)
	l, _ := net.Listen("tcp", ":"+port)

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c, e)
	}

}
