package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error acceptinc con:", err)
		}

		go handleConnection(conn)

	}
}
func handleConnection(conn net.Conn) {

	defer conn.Close()

	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Read Error: %v", err)
		return
	}

	ackMsg := strings.ToUpper(strings.TrimSpace(message))
	response := fmt.Sprintf("ACK: %s\n", ackMsg)
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Printf("Server Write Error: %v", err)
	}

}
