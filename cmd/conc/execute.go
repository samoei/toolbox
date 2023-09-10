package conc

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "concurency",
	Short: "Starts the backoffice reverse proxy",
	Run:   run,
}

func run(_ *cobra.Command, _ []string) {

	serv := newServer()
	go serv.startServer()

	for i := 1; i < 100; i++ {
		msg := fmt.Sprintf("Request from Client %d", i)
		serv.sendMessage(msg)
	}
	time.Sleep(time.Second * 5)
	serv.stopServer()

}

type Server struct {
	stop chan struct{}
	data chan string
}

func newServer() *Server {
	return &Server{
		stop: make(chan struct{}),
		data: make(chan string, 128),
	}
}

func (s *Server) startServer() {
	fmt.Println("Sever starting...")
	s.listenAndServe()
	fmt.Println("Server started and waiting for requests...")
}

func (s *Server) listenAndServe() {

outerLoop:
	for {
		select {
		case <-s.stop:
			break outerLoop
		default:
			time.Sleep(time.Second * 3)
			dat := <-s.data
			s.handleData(dat)
		}
	}
}

func (s *Server) handleData(data string) {
	fmt.Println("Data Received: ", data)
}

func (s *Server) sendMessage(msg string) {
	s.data <- msg
}

func (s *Server) stopServer() {
	fmt.Println("Server shutting down gracefully")
	close(s.stop)
}
