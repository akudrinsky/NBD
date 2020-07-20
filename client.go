package nbd

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

const network = "192.168.0.44"

func LaunchClient(socked_type Socket_type, port string) {
	var conn net.Conn
	if socked_type == Tcp {
		const adress = network + ":" + NBD_DEFAULT_PORT
		var err error
		conn, err = net.Dial(string(Tcp), adress)
		if err != nil {
			fmt.Printf("Error trying to connect with server (%s) %v", adress, err)
			return
		}
	} else {
		fmt.Printf("Non-tcp is NOT ready yet.")
		return
	}

	const maxBufSize uint = 16
	buf := make([]byte, maxBufSize)
	reader := bufio.NewReader(conn)

	if _, err := io.ReadFull(reader, buf); err != nil {
		fmt.Println("Error reading from connection: ", []byte(buf))
		return
	}

	if string(buf) == string(NBDMAGIC)+string(IHAVEOPT) {

	}
}
