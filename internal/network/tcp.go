package network

import (
	"exercise_db/internal/configure"
	"exercise_db/internal/utils"
	"fmt"
	"net"
)

type Networking interface {
	query(sql string) string
}

type Tcp struct {
	lister net.Listener
	config *configure.Config
}

func New(config *configure.Config) *Tcp {
	tcp := &Tcp{}
	listen, err := net.Listen("tcp", config.Network.ToAddress())
	utils.Crash(err)
	tcp.lister = listen
	return tcp
}

func (tcp *Tcp) Start() {
	for {
		conn, err := tcp.lister.Accept()
		if err != nil {
		} else {
			go tcp.Handle(conn)
		}
	}
}

func (tcp *Tcp) Handle(conn net.Conn) {
	defer func() { _ = conn.Close() }()
	buf := make([]byte, 1024)
	protocol := Protocol{}
	used := 0
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		used = 0
		for used < n {
			used = protocol.Write(buf[used:n])
			if protocol.IsFinish() {
				sql := protocol.ToQuery()
				fmt.Println(sql)
			}
		}

	}
}
