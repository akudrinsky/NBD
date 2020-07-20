package nbd

/* List of constants that are needed for NBD protocol */

// NBD magic numbers
const (
	NBDMAGIC                   = 0x4e42444d41474943
	IHAVEOPT                   = 0x49484156454F5054
	REQUEST                    = 0x25609513
	NBD_SIMPLE_REPLY_MAGIC     = 0x67446698
	NBD_STRUCTURED_REPLY_MAGIC = 0x668e33ef
)

// Requests
const (
	NBD_CMD_READ = 0
	NBD_CMD_DISC = 2
)

// Handshake options
const (
	NBD_OPT_GO               = 7
	NBD_OPT_STRUCTURED_REPLY = 8
)

// NBD default port
const (
	NBD_DEFAULT_PORT = "10809"
)

type Socket_type string

// socket type (default: tcp)
const (
	Tcp  Socket_type = "tcp"
	Unix Socket_type = "unix"
)

func int32toArr(num int) []byte {
	const bytesInInt = 4
	const mask = 0b1111111
	const bitsInByte = 8

	answer := make([]byte, bytesInInt) // Maybe problem here (in C++ however)
	for i := 0; i < bytesInInt; i++ {
		answer[i] = byte(mask & num)
		num >>= bitsInByte
	}

	return answer
}

func int64toArr(num int64) []byte {
	const bytesInInt = 8
	const mask = 0b1111111
	const bitsInByte = 8

	answer := make([]byte, bytesInInt) // Maybe problem here (in C++ however)
	for i := 0; i < bytesInInt; i++ {
		answer[i] = byte(mask & num)
		num >>= bitsInByte
	}

	return answer
}
