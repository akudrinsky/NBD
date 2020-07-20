package nbd

/* List of structures (mainly requests and replies),
that are needed for NBD protocol */

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

type NBDrequest struct {
	Magic,
	CommandFlags,
	Type,
	Handle,
	Offset,
	Length []byte
	// TODO: And data to write for NBD_CMD_WRITE
}

// TODO: make it prettier (how?)
func (req *NBDrequest) Read_request(conn net.Conn) {
	reader := bufio.NewReader(conn)

	if _, err := io.ReadAtLeast(reader, req.Magic, 4); err != nil {
		fmt.Println("Error reading request (Magic field) from connection")
		return
	}

	if _, err := io.ReadAtLeast(reader, req.CommandFlags, 2); err != nil {
		fmt.Println("Error reading request (CommandFlags field) from connection")
		return
	}

	if _, err := io.ReadAtLeast(reader, req.Type, 2); err != nil {
		fmt.Println("Error reading request (Type field) from connection")
		return
	}

	if _, err := io.ReadAtLeast(reader, req.Handle, 8); err != nil {
		fmt.Println("Error reading request (Handle field) from connection")
		return
	}

	if _, err := io.ReadAtLeast(reader, req.Offset, 8); err != nil {
		fmt.Println("Error reading request (Offset field) from connection")
		return
	}

	if _, err := io.ReadAtLeast(reader, req.Length, 4); err != nil {
		fmt.Println("Error reading request (Length field) from connection")
		return
	}
}

type NBD_simple_reply struct {
	Magic,
	Error,
	Handle,
	Data []byte
}

func (reply *NBD_simple_reply) Send(conn net.Conn) {
	conn.Write(reply.Magic)
	conn.Write(reply.Error)
	conn.Write(reply.Handle)
	conn.Write(reply.Data)
}

type NBD_structured_reply struct {
	Magic,
	Flags,
	Type,
	Handle []byte
	Length int
	Data   []byte
}

func (reply *NBD_structured_reply) Send(conn net.Conn) {
	// TODO: ...
}
