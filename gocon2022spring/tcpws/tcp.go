package tcpws

import (
	"bufio"
	"io"
	"log"
	"net"
)

type HijackedConn struct {
	rawConn net.Conn
	reader  *bufio.Reader
}

var _ io.ReadWriter = (*HijackedConn)(nil)

func NewHijackedConn(conn net.Conn, reader *bufio.Reader) *HijackedConn {
	return &HijackedConn{
		rawConn: conn,
		reader:  reader,
	}
}

// Read implements io.Reader interface.
func (c *HijackedConn) Read(b []byte) (n int, err error) {
	if c.reader != nil && c.reader.Buffered() > 0 {
		n, err = io.MultiReader(c.reader, c.rawConn).Read(b)

		log.Printf("read from the TCP connection using the buffer(length: %d): %q", n, b)
		return
	}
	n, err = c.rawConn.Read(b)
	if err != nil {
		return
	}

	log.Printf("read from the TCP connection(length: %d): %q", n, b)
	return
}

// Write implements io.Writer interface.
func (c *HijackedConn) Write(b []byte) (int, error) {
	log.Printf("write to the TCP connection(length: %d): %q", len(b), b)

	return c.rawConn.Write(b)
}
