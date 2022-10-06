package main

import (
	"fmt"
	"github.com/kaiaverkvist/kondo"
	"github.com/kaiaverkvist/kondo/examples"
	"github.com/labstack/gommon/log"
	"github.com/olahol/melody"
	"time"
)

func main() {
	s := kondo.NewServer(":8991", "/ws")

	kondo.OnConnect(func(sender *melody.Session) {
		log.Infof("Client (%s) connected to the server!", sender.RemoteAddr())

		welcome := examples.WelcomeMessage{
			Greeting: fmt.Sprintf("(%s) has joined!", sender.RemoteAddr()),
			CurTime:  time.Now(),
		}
		s.Broadcast(welcome)
	})

	kondo.OnDisconnect(func(sender *melody.Session) {
		log.Infof("Client (%s) has disconnected from the server!", sender.RemoteAddr())
	})

	kondo.On[examples.WelcomeMessage](func(sender *melody.Session, message examples.WelcomeMessage) {
		log.Infof("Received greeting from client (%s): %s", sender.RemoteAddr(), message.Greeting)
	})

	s.Init()
}
