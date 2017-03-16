package nucleo

import (
	"fmt"
	"github.com/nats-io/gnatsd/logger"
	"net"
	"strconv"
	"sync"
)

type Module interface {
	OnStart()
	OnMessage()
	OnDestroy()
}

type Skeleton struct {
}

var (
	logMutex sync.Mutex
	log      *logger.Logger
)

//~ var log = struct {}{}

type Server struct {
	Host string
	Port int
}

func setLogger(newLogger *logger.Logger) {
	logMutex.Lock()
	log = newLogger
	logMutex.Unlock()
}

func NewServer() *Server {
	return &Server{
		Host: "0.0.0.0",
		Port: 4333,
	}
}

func (svr *Server) StartServe() {
	setLogger(logger.NewStdLogger(true, true, true, true, true))
	log.Tracef("Starting server [%s]", strconv.Itoa(svr.Port))
	fmt.Printf("Starting server [%d]\n", svr.Port)
	hp := net.JoinHostPort(svr.Host, strconv.Itoa(svr.Port))
	listener, err := net.Listen("tcp", hp)
	if err != nil {
		log.Errorf("Cannot create listening socket")
	}
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Errorf("Error accecepting %v", err)
			}
			//log.Tracef("New connection:%v", conn.RemoteAddr())
			//log.Debugf("New connection:%v\n", conn.RemoteAddr())
			NewClient(conn).StartServe()
		}
	}()
}
