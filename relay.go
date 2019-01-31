package main

import "net"
import "fmt"
import "bufio"
import "os"

func relayMessage() {

	var echoServerURL = "localhost:8080"
	conn, _ := net.Dial("tcp", echoServerURL)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send
		fmt.Fprintf(conn, text+"\n")
		// listen & return message
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
