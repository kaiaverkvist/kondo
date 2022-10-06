# kondo
Kondo is a wrapper for Melody which adds strongly typed routing.

## 💯 Usage

*This is an example snippet which can be found in the /examples folder.*

### Server
> This sets up a basic server which will broadcast a message to all sessions when someone connects.
```go
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
```

### Client
> This sets up a basic client which will log to console whenever the server sends a greeting.
```go
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
```

## Contributions
* Contributions are welcome.