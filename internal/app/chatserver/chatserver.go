// Package chatserver as main server app
package chatserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"web-chat-on-golang/dbstorage"
)

// ChatServer structure
type ChatServer struct {
	config   *Config
	router   *mux.Router
	database *dbstorage.DbStorage
}

// New ChatServer init
func New(config *Config) *ChatServer {
	return &ChatServer{
		config: config,
		router: mux.NewRouter(),
	}
}

// Start the server
func (s *ChatServer) Start() error {
	s.configureRouter()
	if err := s.configureDatabase(); err != nil {
		return err
	}

	fmt.Println("Chat server started")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// configureRouter with different paths
func (s *ChatServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handlerHello())
}

// configureDatabase to operate
func (s *ChatServer) configureDatabase() error {
	d := dbstorage.New(s.config.Database)
	if err := d.Open(); err != nil {
		return err
	}
	s.database = d
	return nil
}

// handlerHello to respond on certain request
func (s *ChatServer) handlerHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}
