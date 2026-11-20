package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("listening on :9000")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	log.Printf("client connected: %s", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("received: %q", line)

		fmt.Fprintf(conn, "echo: %s\n", line)
	}

	log.Printf("client disconnected: %s", conn.RemoteAddr())
}
