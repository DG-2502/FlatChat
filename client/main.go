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

	done := make(chan struct{})

	go func() {
		defer close(done)
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	go func() {
		stdin := bufio.NewScanner(os.Stdin)
		for stdin.Scan() {
			fmt.Fprintln(conn, stdin.Text())
		}
	}()

	<-done
	fmt.Println("[disconnected]")
}
