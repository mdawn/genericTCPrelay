package main

import (
	"bufio"
	"flag"
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
	port := flag.String("p", "", "port")
	flag.Usage = func() {
		flag.PrintDefaults()
	}
	flag.Parse()

	if *port == "" {
		flag.Usage()
		os.Exit(0)
	}

	l, err := net.Listen("tcp", ":"+*port)
	handleError(err, "Listen()")

	defer l.Close()
	fmt.Println("Listening on Port", *port)
	conn, err := l.Accept()
	handleError(err, "Accept()")

	defer conn.Close()
	fmt.Printf("Connected to : %v\n", conn.RemoteAddr())
	reader := bufio.NewReader(conn)
	for {
		_, err = fmt.Fprintf(conn, "Exit using STOP \r\n")
		handleError(err, "first Fprintf()")

		str, err := readln(reader)
		handleError(err, "readln()")
		if str == "stop" {
			os.Exit(0)
		}
		fmt.Println("Input:" + str)
		_, err = fmt.Fprintf(conn, "Input: %s\r\n", str)
		handleError(err, "second Fprintf()")
	}
}
