package websocket

import (
	"net/http"
	"time"
)

type option struct {
	header       http.Header
	pingInterval time.Duration
	pingHandler  func(appData string) error
	pongHandler  func(appData string) error
	readDeadline time.Duration
	closeHandler func(code int, text string) error
}

var defaultOption = &option{
	pingInterval: time.Second * 30,
}

type Option func(*option)

func WithHeader(header http.Header) Option {
	return func(o *option) {
		o.header = header
	}
}

// WithPingInterval 发送ping时间间隔
// default 30s
func WithPingInterval(pingInterval time.Duration) Option {
	return func(o *option) {
		o.pingInterval = pingInterval
	}
}

func WithPingHandler(h func(appData string) error) Option {
	return func(o *option) {
		o.pingHandler = h
	}
}
func WithPongHandler(h func(appData string) error) Option {
	return func(o *option) {
		o.pongHandler = h
	}
}

func WithReadDeadline(t time.Duration) Option {
	return func(o *option) {
		o.readDeadline = t
	}
}

func WithCloseHandler(h func(code int, text string) error) Option {
	return func(o *option) {
		o.closeHandler = h
	}
}

const (
	ErrorMessage = -1
	// TextMessage denotes a text data message. The text message payload is
	// interpreted as UTF-8 encoded text data.
	TextMessage = 1

	// BinaryMessage denotes a binary data message.
	BinaryMessage = 2

	// CloseMessage denotes a close control message. The optional message
	// payload contains a numeric code and text. Use the FormatCloseMessage
	// function to format a close message payload.
	CloseMessage = 8

	// PingMessage denotes a ping control message. The optional message payload
	// is UTF-8 encoded text.
	PingMessage = 9

	// PongMessage denotes a pong control message. The optional message payload
	// is UTF-8 encoded text.
	PongMessage = 10
)
