package websocket

import (
	"context"
	"errors"
	"sync"
	"time"

	ws "github.com/gorilla/websocket"
)

type Client struct {
	conn      *ws.Conn
	option    *option
	readMsgCh chan *Msg
	mu        sync.Mutex
}

type Msg struct {
	MsgType int
	Msg     []byte
	Err     error
}

func NewClient(path string, opts ...Option) (*Client, error) {
	var client = &Client{
		readMsgCh: make(chan *Msg, 100),
	}
	client.option = defaultOption
	for _, v := range opts {
		v(client.option)
	}

	conn, _, err := ws.DefaultDialer.Dial(path, client.option.header)
	if err != nil {
		return nil, err
	}

	client.conn = conn
	client.setConnOption()

	return client, nil
}
func (c *Client) setConnOption() {
	if c.option.pingHandler != nil {
		c.conn.SetPingHandler(c.option.pingHandler)
	}
	if c.option.pongHandler != nil {
		c.conn.SetPongHandler(c.option.pongHandler)
	}
	if c.option.closeHandler != nil {
		c.conn.SetCloseHandler(c.option.closeHandler)
	}

}

func (c *Client) readMsg() {
	for {
		msgType, data, err := c.conn.ReadMessage()
		if msgType == CloseMessage || msgType == ErrorMessage || err != nil {
			return
		}
		c.readMsgCh <- &Msg{MsgType: msgType, Msg: data, Err: err}
	}
}

func (c *Client) ReadMessage(ctx context.Context, callback func(*Msg, error)) error {
	if c.conn == nil {
		return errors.New("client is not init")
	}
	go c.readMsg()

	go func() {
		ticker := time.NewTicker(c.option.pingInterval)
		defer func() {
			ticker.Stop()
		}()
		for {
			select {
			case <-ctx.Done():
				err := c.conn.WriteMessage(CloseMessage, nil)
				if err != nil {
					callback(&Msg{MsgType: CloseMessage}, err)
				}
				return
			case <-ticker.C:
				err := c.conn.WriteMessage(PingMessage, nil)
				if err != nil {
					callback(nil, err)
					return
				}
			case msg := <-c.readMsgCh:
				callback(msg, nil)
			}
		}
	}()
	return nil
}

func (c *Client) WriteMessage(msgType int, msg []byte) error {
	if c.conn == nil {
		return errors.New("client is not init")
	}
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()

	return c.conn.WriteMessage(msgType, msg)
}
