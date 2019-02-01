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

func readln2(r *bufio.Reader) (string, error) {
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
	port := "8083"

	l, err := net.Listen("tcp", ":"+port)
	handleError2(err, "Listen()")

	defer l.Close()
	fmt.Println("Listening on Port", port)
	conn, err := l.Accept()
	handleError2(err, "Accept()")

	defer conn.Close()
	fmt.Printf("Connected to : %v\n", conn.RemoteAddr())
	reader := bufio.NewReader(conn)
	for {

		str, err := readln2(reader)
		handleError2(err, "readln()")
		if str == "stop" {
			os.Exit(0)
		}
		_, err = fmt.Fprintf(conn, "%s\r\n", str)
		handleError2(err, "Fprintf()")
	}

}
