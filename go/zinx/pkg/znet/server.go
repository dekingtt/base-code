package znet

import (
	"fmt"
	"net"
	"time"
	"zinx/pkg/ziface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server listenner at IP %s, Port: %d", s.IP, s.Port)

	go func() {
		// get a tcp addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		// listenning
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}

		// start to lisenning
		fmt.Println("start Zinx server ", s.Name, " success, now listenning...")

		for {
			// wait for connection
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			// TODO server.start()

			// TODO server.start()

			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buffer error", err)
						continue
					}

					// echo back
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buffer error", err)
						continue
					}
				}
			}()
		}

	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server, name ", s.Name)
	// TODO Server.stop()
}

func (s *Server) Serve() {
	s.Start()

	// TODO Server.Serve()

	for {
		time.Sleep(10 * time.Second)
	}
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7777,
	}

	return s
}
