package NBD

/* List of constants that are needed for NBD protocol */

// NBD magic numbers
const (
	NBD_MAGIC                  = 0x4e42444d41474943
	NBD_REQUEST_MAGIC          = 0x25609513
	NBD_REPLY_MAGIC            = 0x67446698
	NBD_CLISERV_MAGIC          = 0x00420281861253
	NBD_OPTS_MAGIC             = 0x49484156454F5054
	NBD_REP_MAGIC              = 0x3e889045565a9
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