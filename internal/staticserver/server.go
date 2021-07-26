package staticserver

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/joeyscat/oook/internal/util"
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
	http.HandleFunc("/upload", s.upload())

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

func (s *StaticServer) upload() func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		uploadDir := path.Join(s.directory, "upload")

		if !util.DirExists(uploadDir) {
			err := os.Mkdir(uploadDir, 0755)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer file.Close()

		fmt.Fprintf(w, "%v", handler.Header)

		filepath := path.Join(uploadDir, handler.Filename)
		uploadDest, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer uploadDest.Close()
		io.Copy(uploadDest, file)
	}
}
