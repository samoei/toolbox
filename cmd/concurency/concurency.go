package concurency

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var ConcurencyCommand = &cobra.Command{
	Use:   "concurency",
	Short: "Starts the backoffice reverse proxy",
	Run:   runConcurency,
}

func runConcurency(_ *cobra.Command, _ []string) {

	//iknitiase the server
	serv := newServer()

	//run the server in another thread
	go serv.startServer()

	//Send messages to the server
	for i := 1; i < 100; i++ {
		msg := fmt.Sprintf("Request from Client %d", i)
		serv.sendMessage(msg)
	}

	// stop the server after 10 seconds
	time.Sleep(time.Second * 10)
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
		// This is the done channel
		case <-s.stop:
			break outerLoop
		//Do this untill we rececive a signal on the done channel
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
