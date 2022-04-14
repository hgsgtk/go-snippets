package tcpws

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Method int

const (
	_ Method = iota
	Handshake
	Completed
	Communication
)

type Protocol struct {
	Method Method
	Dest   string // Use for Handshake
	Data   []byte // Use for Communication
}

type WrapWSConn struct {
	RawConn *websocket.Conn
}

func (c *WrapWSConn) ReadHandshakeRequest() (string, error) {
	_, rb, err := c.RawConn.ReadMessage()
	if err != nil {
		return "", err
	}
	var dc Protocol
	if err := json.Unmarshal(rb, &dc); err != nil {
		return "", err
	}

	if dc.Method != Handshake {
		return "", fmt.Errorf("unexpected method: %d", dc.Method)
	}
	return dc.Dest, nil
}

func (c *WrapWSConn) WriteHandshakeRequest(dest string) error {
	proto := Protocol{
		Method: Handshake,
		Dest:   dest,
	}
	db, err := json.Marshal(proto)
	if err != nil {
		return err
	}

	return c.RawConn.WriteMessage(websocket.TextMessage, db)
}

func (c *WrapWSConn) WriteHandshakeCompleted() error {
	proto := Protocol{
		Method: Completed,
	}
	db, err := json.Marshal(proto)
	if err != nil {
		return err
	}

	return c.RawConn.WriteMessage(websocket.TextMessage, db)
}

func (c *WrapWSConn) IsHandshaked() error {
	_, rb, err := c.RawConn.ReadMessage()
	if err != nil {
		return err
	}
	var dc Protocol
	if err := json.Unmarshal(rb, &dc); err != nil {
		return err
	}

	if dc.Method != Completed {
		return errors.New("unexpected method")
	}

	return nil
}

func (c *WrapWSConn) Write(p []byte) (n int, err error) {
	proto := Protocol{
		Method: Communication,
		Data:   p,
	}
	db, err := json.Marshal(proto)
	if err != nil {
		return 0, err
	}

	if err := c.RawConn.WriteMessage(websocket.BinaryMessage, db); err != nil {
		return 0, err
	}
	log.Printf("write to the WebSocket connection: %s", p)
	return len(db), nil
}

func (c *WrapWSConn) Read(p []byte) (n int, err error) {
	_, rb, err := c.RawConn.ReadMessage()
	if err != nil {
		return 0, err
	}
	var dc Protocol
	if err := json.Unmarshal(rb, &dc); err != nil {
		return 0, err
	}

	if dc.Method != Communication {
		return 0, fmt.Errorf("unexpected method: %d", dc.Method)
	}

	p = dc.Data
	log.Printf("read from the WebSocket connection: %s", p)
	return len(p), nil
}
