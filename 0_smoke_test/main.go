package main

import (
	"io"
	"log"
	"net"

	"github.com/hhow09/protohackers/server"
)

// problem: https://protohackers.com/problem/0
// ref: https://go.dev/src/net/example_test.go

func main() {
	s := server.New("tcp", "localhost:8080")
	s.Handle(handle)
}

func handle(conn net.Conn) {
	defer conn.Close()
	// use ReadAll because of unknown buffer size
	buf, err := io.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Received: %s", buf)
	_, err = conn.Write(buf)
	if err != nil {
		log.Fatal(err)
	}
}
