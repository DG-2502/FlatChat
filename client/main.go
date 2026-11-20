package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	serverReader := bufio.NewReader(conn)
	stdin := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !stdin.Scan() {
			return
		}

		fmt.Fprintf(conn, "%s\n", stdin.Text())

		reply, err := serverReader.ReadString('\n')
		if err != nil {
			log.Println("read error:", err)
			return
		}
		fmt.Print(reply)
	}
}
