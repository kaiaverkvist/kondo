package kondo

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/olahol/melody"
)

// Server is the main component of the NECS server.
// Instantiate this to host an ECS client.
type Server struct {
	e *echo.Echo
	m *melody.Melody

	path        string
	bindAddress string
}

// NewServer instantiates a NECS server.
// Specify the `bindAddress` a full address with port.
// Example: `localhost:8080` or `0.0.0.0:8108`.
// .. `:8080` also works.
func NewServer(bindAddress string, path string) Server {
	return Server{
		e: echo.New(),
		m: melody.New(),

		path:        path,
		bindAddress: bindAddress,
	}
}

func (srv *Server) Init() {
	srv.e.HideBanner = true
	//e.Use(middleware.Logger())
	srv.e.Use(middleware.Recover())

	srv.e.GET(srv.path, func(c echo.Context) error {
		return srv.m.HandleRequest(c.Response().Writer, c.Request())
	})

	srv.m.HandleConnect(func(session *melody.Session) {
		callConnect(session)
	})

	srv.m.HandleDisconnect(func(session *melody.Session) {
		callDisconnect(session)
	})

	srv.m.HandleMessageBinary(func(s *melody.Session, msg []byte) {
		nm, err := parseIncoming(msg)
		if err != nil {
			log.Errorf("Unable to process message with length %d, and sender (%s): '%s'", len(msg), s.RemoteAddr(), err)
		}

		err = processMessage(s, nm)
		if err != nil {
			log.Errorf("Unable to process message with length %d, and sender (%s): '%s'", len(msg), s.RemoteAddr(), err)
		}
	})

	srv.e.Logger.Fatal(srv.e.Start(srv.bindAddress))
}

// Broadcast sends a message to all connected sessions.
func (srv *Server) Broadcast(data any) {
	_ = srv.m.BroadcastBinary(message(data))
}

// SendFiltered sends the broadcast exclusively to sessions which the filter function returns true on.
func (srv *Server) SendFiltered(data any, filter func(session *melody.Session) bool) {
	_ = srv.m.BroadcastBinaryFilter(message(data), filter)
}

// BroadcastOthers sends a message to all connected sessions except session s.
func (srv *Server) BroadcastOthers(data any, s *melody.Session) {
	_ = srv.m.BroadcastBinaryOthers(message(data), s)
}

// SendTo sends a message to a specific session.
func (srv *Server) SendTo(data any, s *melody.Session) {
	_ = s.WriteBinary(message(data))
}

func (srv *Server) GetSessions() ([]*melody.Session, error) {
	return srv.m.Sessions()
}
