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

func readln(r *bufio.Reader) (string, error) {
	input, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = input[:len(input)-1]
	if input[len(input)-1] == '\r' {
		input = input[:len(input)-1]
	}
	return input, nil
}

func main() {
	var relayServer = "localhost:8080"

	conn, _ := net.Dial("tcp", relayServer)
	reader := bufio.NewReader(conn)

	for {

		str, err := readln(reader)
		handleError(err, "readln()")
		if str == "stop" {
			os.Exit(0)
		}
		_, err = fmt.Fprintf(conn, "%s\r\n", str)
		handleError(err, "Fprintf()")
	}
}
