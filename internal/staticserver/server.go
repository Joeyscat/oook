package staticserver

import (
	"fmt"
	"net/http"
)

type StaticServer struct {
	directory string
	port      uint
}

func NewStaticServer(directory string, port uint) *StaticServer {
	return &StaticServer{
		directory: directory,
		port:      port,
	}
}

func (s *StaticServer) Run() error {
	http.Handle("/", http.FileServer(http.Dir(s.directory)))

	fmt.Printf("Static Server Running on http://127.0.0.1:%d/\n", s.port)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
}
