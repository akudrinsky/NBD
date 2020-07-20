// Package nbd is made for launching NBD protocol
// TODO: write docs
package nbd // TODO: test package

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

// LaunchServer TODO write docs
func LaunchServer(sockedType Socket_type, port string) {
	fmt.Println("Launching server...")
	ln, err := net.Listen(string(sockedType), port)

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

		if successfullHandshake(conn) {
			transmission(conn)
		}
	}
}

func successfullHandshake(conn net.Conn) bool {
	conn.Write(int64toArr(NBDMAGIC))
	conn.Write(int64toArr(IHAVEOPT))

	conn.Write([]byte{0b1, 0}) // 0b1 - zero-th bit is set to 1 because fixed newstyle

	reader := bufio.NewReader(conn)
	const bufSize uint = 4
	clientOptions := make([]byte, bufSize)

	if _, err := io.ReadFull(reader, clientOptions); err != nil {
		fmt.Println("Error reading from connection: ", clientOptions)
		conn.Close()
		return false
	}

	if clientOptions[0] != 1 ||
		clientOptions[1] != 0 ||
		clientOptions[2] != 0 ||
		clientOptions[3] != 0 {
		fmt.Println("Disconnection due to unknown client options (need 100...00): ", clientOptions)
		conn.Close()
		return false
	}

	return true
}

func transmission(conn net.Conn) {
	var request NBDrequest
	request.Read_request(conn)

	toInt := func(array []byte) int {
		return int(array[1])<<1 + int(array[0])
	}

	switch toInt(request.Type) {
	case NBD_CMD_READ:
		// TODO: structured reply
		simpleReply := NBD_simple_reply{
			int32toArr(NBD_SIMPLE_REPLY_MAGIC),
			[]byte{0, 0, 0, 0}, /* TODO: errors */
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
