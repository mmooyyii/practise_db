package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"os"
)

var (
	host string
	port int

	username string
	password string
)

type Session struct {
	data  []byte
	inStr bool
}

func NewSession() *Session {
	return &Session{
		data:  make([]byte, 0),
		inStr: false,
	}
}

func (s *Session) readline() error {
	reader := bufio.NewReader(os.Stdin)
	data, _, err := reader.ReadLine()
	if err != nil {
		s.data = []byte{}
		return err
	}
	n := bytes.Count(data, []byte(`'`))
	s.inStr = (n%2 == 0) == s.inStr
	s.data = append(s.data, data...)
	return nil
}

func (s *Session) prefix() string {
	if s.inStr {
		return "'>"
	}
	return "->"
}

func (s *Session) exec() []string {
	return []string{"123123"}
}

func main() {
	flag.StringVar(&host, "host", "127.0.0.1", "connect to database")
	flag.IntVar(&port, "port", 12345, "connect to database")
	flag.StringVar(&username, "username", "root", "connect to database")
	flag.StringVar(&password, "password", "", "connect to database")
	flag.Parse()
	session := NewSession()
	var err error
	for {
		fmt.Print(session.prefix())
		err = session.readline()
		if err != nil {
			err = nil
			fmt.Println(err.Error())
		}
		exec := session.exec()
		for _, v := range exec {
			fmt.Println(v)
		}
	}
}

func makeRequest(body string) ([]byte, error) {
	magicHead := uint16(0xFFFF)
	size := int32(len(body))
	b := &bytes.Buffer{}
	err := binary.Write(b, binary.LittleEndian, magicHead)
	if err != nil {
		return nil, err
	}
	err = binary.Write(b, binary.LittleEndian, size)
	if err != nil {
		return nil, err
	}
	err = binary.Write(b, binary.LittleEndian, []byte(body))
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
