package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)

type server struct {
	mu      sync.Mutex
	clients map[net.Conn]struct{}
}

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("listening on :9000")

	s := &server{clients: make(map[net.Conn]struct{})}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}
		go s.handle(conn)
	}
}

func (s *server) handle(conn net.Conn) {
	defer func() {
		s.mu.Lock()
		delete(s.clients, conn)
		s.mu.Unlock()
		conn.Close()
	}()

	s.mu.Lock()
	s.clients[conn] = struct{}{}
	s.mu.Unlock()

	log.Printf("connected: %s", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("received from %s: %q", conn.RemoteAddr(), line)
		s.broadcast(line)
	}

	log.Printf("disconnected: %s", conn.RemoteAddr())
}

func (s *server) broadcast(line string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for conn := range s.clients {
		fmt.Fprintf(conn, "%s\n", line)
	}
}
