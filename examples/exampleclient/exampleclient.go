package main

import (
	"github.com/kaiaverkvist/kondo"
	"github.com/kaiaverkvist/kondo/examples"
	"github.com/labstack/gommon/log"
	"github.com/olahol/melody"
	"net/url"
)

func main() {
	c := kondo.NewClient(url.URL{Scheme: "ws", Host: "localhost:8991", Path: "/ws"})

	kondo.OnConnect(func(sender *melody.Session) {
		log.Infof("I have connected to the server!")
	})

	kondo.OnDisconnect(func(sender *melody.Session) {
		log.Infof("I have disconnected to the server!")
	})

	kondo.On[examples.WelcomeMessage](func(sender *melody.Session, message examples.WelcomeMessage) {
		log.Infof("Received greeting from server: %s", message.Greeting)
	})

	c.Init()
}
