package staticserver

type StaticServer struct {
	Path string
	Port uint
}

func NewStaticServer(path string, port uint) *StaticServer {
	return &StaticServer{
		Path: path,
		Port: port,
	}
}

func (*StaticServer) Run() error {

	return nil
}
