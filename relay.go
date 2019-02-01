package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func handleError2(err error, msg string) {
	if err != nil {
		fmt.Printf("Error: %s : %s\n", msg, err.Error())
		os.Exit(1)
	}
}

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

	// server port
	port := "8080"

	l, err := net.Listen("tcp", ":"+port)
	s, _ := l.Accept()

	handleError2(err, "Listen()")

	defer l.Close()
	fmt.Println("Listening on Port", port)
	handleError2(err, "Accept()")

	defer s.Close()
	fmt.Printf("Connected to : %v\n", s.RemoteAddr())
	bufio.NewReader(s)

	// client port
	port2 := "8081"

	t, err := net.Listen("tcp", ":"+port2)

	handleError2(err, "Listen()")

	defer t.Close()
	fmt.Println("Listening on Port", port2)
	handleError2(err, "Accept()")

	for {
		c, err := t.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c, s)
	}

}
