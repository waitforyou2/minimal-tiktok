package config

type Server struct {
	Port int
}

var DefaultServer = &Server{
	Port: 8080,
}
