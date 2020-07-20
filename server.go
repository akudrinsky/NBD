// Package nbd is made for launching NBD protocol
// TODO: write docs
package nbd // TODO: test package

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

// Launch_server TODO write docs
func LaunchServer(socked_type Socket_type, port string) {
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
			transmission(conn)
		}
	}
}

func successfull_handshake(conn net.Conn) bool {
	fmt.Fprintf(conn, string(NBDMAGIC)) // correct?
	fmt.Fprintf(conn, string(IHAVEOPT))

	// TODO: send handshake flags (?)

	// TODO: accept client flags (close connection and return false if does not recognize)
	reader := bufio.NewReader(conn)
	const bufSize uint = 4
	buf := make([]byte, bufSize)

	if _, err := io.ReadFull(reader, buf); err != nil {
		fmt.Println("Error reading from connection: ", []byte(buf))
		return false
	}

	return true
}

func transmission(conn net.Conn) {
	var request NBDrequest
	request.Read_request(conn)

	toInt := func(array []byte) int { return int(array[1])<<1 + int(array[0]) }

	switch toInt(request.Type) {
	case NBD_CMD_READ:
		// TODO: structured reply
		simpleReply := NBD_simple_reply{
			toArr(NBD_SIMPLE_REPLY_MAGIC),
			[]byte{0}, /* TODO: errors */
			request.Handle,
			[]byte("Hello world!"), /* TODO: data to send */
		}

		simpleReply.Send(conn)

	case NBD_CMD_DISC:
		// Handle all outstanding requests!
		conn.Close()
		// Close TCP session (???)
	}

}
