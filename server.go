package NBD

import (
	"bufio"
	"fmt"
	"net"
	"io"
)

func Launch_server(socked_type Socket_type, port string) {
	fmt.Println("Launching server...")
	ln, err := net.Listen(string(socked_type), port)

	if err != nil {
		fmt.Println("Error in listening port.")
		ln.Close()
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection %v", err)
			continue
		}
		fmt.Printf("Accepted connection from %v", conn.RemoteAddr())

		if successfull_handshake(conn) {

		}
	}
}

func successfull_handshake(conn net.Conn) bool {
	fmt.Fprintf(conn, string(NBDMAGIC)) // correct?
	fmt.Fprintf(conn, string(IHAVEOPT))

	// TODO: send handshake flags (?)

	// TODO: accept client flags (close connection and return false if does not recognize)
	reader := bufio.NewReader(conn)
	const buf_size uint = 4
	buf := make([]byte, buf_size)

	if _, err := io.ReadFull(reader, buf); err != nil {
		fmt.Println("Error reading from connection: ", []byte(buf))
		return false
	}

	return true
}

func read_bytes(where []byte, number uint) {
	
}