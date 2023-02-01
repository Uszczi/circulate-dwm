package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type Message struct {
	ID   string
	Data string
}

func send(conn net.Conn) {
	msg := Message{ID: "Yo", Data: "Hello"}
	bin_buf := new(bytes.Buffer)

	gobobj := gob.NewEncoder(bin_buf)
	_ = gobobj.Encode(msg)

	_, _ = conn.Write(bin_buf.Bytes())
}

func recv(conn net.Conn) {
	tmp := make([]byte, 500)
	_, _ = conn.Read(tmp)

	tmpbuff := bytes.NewBuffer(tmp)
	tmpstruct := new(Message)

	gobobjdec := gob.NewDecoder(tmpbuff)
	_ = gobobjdec.Decode(tmpstruct)

	fmt.Println(tmpstruct)
}

func SendCommand() {
	// TODO make ip passable
	conn, _ := net.Dial("tcp", "127.0.0.1:8018")

	send(conn)
	recv(conn)
}

// TODO there has to be my own logger
func logerr(err error) bool {
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			log.Println("read timeout:", err)
		} else if err == io.EOF {
		} else {
			log.Println("read error:", err)
		}
		return true
	}
	return false
}

func read(conn net.Conn) {
	tmp := make([]byte, 500)

	for {
		_, err := conn.Read(tmp)
		if logerr(err) {
			break
		}

		tmpbuff := bytes.NewBuffer(tmp)
		tmpstruct := new(Message)

		gobobj := gob.NewDecoder(tmpbuff)
		_ = gobobj.Decode(tmpstruct)

		fmt.Println(tmpstruct)
		return
	}
}

func resp(conn net.Conn) {
	msg := Message{ID: "Yo", Data: "Hello back"}
	bin_buf := new(bytes.Buffer)

	gobobje := gob.NewEncoder(bin_buf)
	_ = gobobje.Encode(msg)

	_, _ = conn.Write(bin_buf.Bytes())
	conn.Close()
}

func handle(conn net.Conn) {
	timeoutDuration := 2 * time.Second
	fmt.Println("Launching server...")
	_ = conn.SetReadDeadline(time.Now().Add(timeoutDuration))

	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)

	read(conn)
	resp(conn)
}

func RunTcpServer() {
	server, _ := net.Listen("tcp", "localhost:8018")
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Println("Connection error: ", err)
			return
		}
		go handle(conn)
	}
}
