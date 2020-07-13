package NBD

import (
	"fmt"
	"net"
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

		if handshake(conn) {

		}
    }
}

func handshake(conn net.Conn) bool {
	
}