package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
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
		fmt.Println("message from client: " + text)
		// send
		fmt.Fprintf(e, text)
		// listen & return message
		m, _ := bufio.NewReader(e).ReadString('\n')
		fmt.Println("Message from echo server: " + m)
		// passes msg along
		fmt.Fprintf(c, "%s\r\n", m)
	}
}

func multiServ(s net.Conn, p int) {

	defer s.Close()
	fmt.Printf("Connected to: %v\n", s.RemoteAddr())
	bufio.NewReader(s)

	fmt.Fprintf(s, "established relay address: localhost "+opt.port+"\n")

	// client handling

	// client port

	t, err := net.Listen("tcp", ":"+strconv.Itoa(p))

	handleError2(err, "Listen()")

	defer t.Close()
	fmt.Println("Listening for client on Port", p)
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

type options struct {
	port string
}

var opt options

func main() {

	//server handling

	flag.StringVar(&opt.port, "p", os.Getenv("PORT"), "The default port to listen on")
	flag.Parse()

	if opt.port == "" {
		opt.port = "7"
	}

	flag.Parse()

	log.Printf("[INFO] Listening on %s\n", opt.port)

	l, err := net.Listen("tcp", ":"+opt.port)

	p := 8081

	// accept concurrent server & cleint connections
	for {
		s, _ := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go multiServ(s, p)
		p++
	}

}
