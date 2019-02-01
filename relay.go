package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func handleError(err error, msg string) {
	if err != nil {
		fmt.Printf("Error: %s : %s\n", msg, err.Error())
		os.Exit(1)
	}
}

func main() {
	fmt.Println("relay runs")

	var echoServerURL = "localhost:8080"
	e, _ := net.Dial("tcp", echoServerURL)

	port := "8081"

	l, err := net.Listen("tcp", ":"+port)
	handleError(err, "Listen()")

	defer l.Close()
	fmt.Println("Listening on Port", port)
	c, err := l.Accept()
	handleError(err, "Accept()")

	defer c.Close()
	fmt.Printf("Connected to : %v\n", c.RemoteAddr())
	reader := bufio.NewReader(c)

	for {
		text, _ := reader.ReadString('\n')
		// send
		fmt.Fprintf(e, text+"\n")
		// listen & return message
		m, _ := bufio.NewReader(e).ReadString('\n')
		fmt.Print("Message from server: " + m)
		// passes msg along
		_, err = fmt.Fprintf(c, "%s\r\n", m)
		handleError(err, "Fprintf()")
	}

}
