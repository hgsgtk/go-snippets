package tcpws

import (
	"bufio"
	"io"
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
		return io.MultiReader(c.reader, c.rawConn).Read(b)
	}
	return c.rawConn.Read(b)
}

// Write implements io.Writer interface.
func (c *HijackedConn) Write(b []byte) (int, error) {
	return c.rawConn.Write(b)
}
