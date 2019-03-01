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
	var relayServer = "localhost:8090"

	conn, _ := net.Dial("tcp", relayServer)
	reader := bufio.NewReader(conn)

	// read public address from relay

	addr, err := reader.ReadString('\n')
	fmt.Println("read an address string: " + addr)
	if err != nil {
		fmt.Println("there was an error reading from client")
		return
	}

	// handleError(err, "readln()1")
	fmt.Println(addr)

	for {
		// listener
		str, err := reader.ReadString('\n')
		fmt.Println("I read the string: " + str)
		if err != nil {
			fmt.Println("there was an error reading the string: " + str)
			return
		}
		_, err = fmt.Fprintf(conn, "%s\r\n", "pong!") // conn
		handleError(err, "Fprintf()")
	}

}
