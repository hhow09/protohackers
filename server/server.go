package server

import (
	"errors"
	"net"
)

type Server struct {
	network string
	address string
	l       net.Listener
}

func New(network, address string) *Server {
	return &Server{
		network: network,
		address: address,
	}
}

func (s *Server) Handle(h func(conn net.Conn)) error {
	var err error
	s.l, err = net.Listen(s.network, s.address)
	if err != nil {
		return err
	}
	for {
		conn, err := s.l.Accept()
		if err != nil {
			return err
		}
		if conn == nil {
			return errors.New("nil connection")
		}
		go h(conn)
	}
}

func (s *Server) Close() (err error) {
	return s.l.Close()
}
