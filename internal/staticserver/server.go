package staticserver

import (
	"fmt"
	"github.com/joeyscat/oook/internal/util"
	"net/http"
	"os"
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

	ip, err := util.GetOutBoundIP()
	if err != nil {
		ip = "127.0.0.1"
		_, _ = fmt.Fprintf(os.Stderr, "We don't know the IP address of your computer, please get it by yourself.\n%v", err)
	}

	addr := fmt.Sprintf("http://%s:%d", ip, s.port)
	fmt.Printf("Static Server Running on %s\n\n", addr)

	util.RenderString(addr)

	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
}
