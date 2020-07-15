package NBD

/* List of constants that are needed for NBD protocol */

// NBD magic numbers
const (
	NBDMAGIC = 0x4e42444d41474943
	IHAVEOPT = 0x49484156454F5054
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
