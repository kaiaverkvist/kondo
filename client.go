package kondo

import (
	"github.com/labstack/gommon/log"
	"github.com/olahol/melody"
	"net/url"
)

// Client is the instance to use when connecting to the hosted ECS on a Server.
type Client struct {
	m *melody.MelodyClient

	connectionUrl url.URL
}

func NewClient(url url.URL) Client {
	c := Client{
		m: melody.NewClient(url),
	}
	c.connectionUrl = url
	return c
}

func (c *Client) Init() {

	c.m.HandleError(func(err error) {
		log.Info("Encountered error from melody: ", err)
	})

	c.m.HandleMessageBinary(func(msg []byte) {
		nm, err := parseIncoming(msg)
		if err != nil {
			log.Error("Unable to decode message with length: ", len(msg), ": ", err)
		}

		// Since we're clientside, sender will be nil!
		err = processMessage(nil, nm)
		if err != nil {
			log.Error("Unable to process message with length: ", len(msg), ": ", err)
		}
	})

	c.m.HandleConnect(func() {
		callConnect(nil)
	})

	c.m.HandleDisconnect(func() {
		callDisconnect(nil)
	})

	c.m.Connect()
}

// SendToServer sends a packet to the server.
func (c *Client) SendToServer(data any) {
	_ = c.m.SendBinary(message(data))
}
