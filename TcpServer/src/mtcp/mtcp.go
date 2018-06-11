package mtcp

import (
	"errtype"
	"log"
	"net"
)

type TcpServer struct {
	addr     string
	listener net.Listener
}

func (server *TcpServer) Start(addr string, newConnection func(conn net.Conn)) (err error) {
	if addr == "" {
		return errtype.BuildStringError("Add is invaild!")
	}
	server.addr = addr
	if listener, err := net.Listen("tcp", server.addr); err != nil {
		return err
	} else {
		server.listener = listener
	}

	for {
		if conn, err := server.listener.Accept(); err != nil {
			log.Println(err)
		} else {
			log.Println("New COnnection!")
			newConnection(conn)
		}
	}

	return nil
}
