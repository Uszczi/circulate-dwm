package core

import (
	"bytes"
	"circulate/circulate/store"
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

func send(conn net.Conn, message string) {
	msg := Message{ID: "Yo", Data: message}
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

func SendCommand(message ...string) {
	// TODO make ip passable
	var msg string
	if len(message) > 0 {
		msg = message[0]
	} else {
		msg = "default"
	}
	conn, _ := net.Dial("tcp", "127.0.0.1:8018")

	send(conn, msg)
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
		fmt.Println("Prevous active layout", store.GetActiveLayout())
		store.SetActiveLayout(tmpstruct.Data)
		if tmpstruct.Data == "toogle" {
			UseSetHowHide()
		} else if tmpstruct.Data == "switch-to-workspace 1" {
			fmt.Println("ssssswitch-to-workspace 1")
		} else if tmpstruct.Data == "switch-to-workspace 2" {
			fmt.Println("ssssswitch-to-workspace 2")
		} else if tmpstruct.Data == "move-to-workspace 1" {
			fmt.Println("ssssmove-to-workspace 1")
		} else if tmpstruct.Data == "move-to-workspace 2" {
			fmt.Println("ssssmove-to-workspace 2")
		} else if tmpstruct.Data == "debug-workspace" {
			fmt.Println("ssssdebug-workspace")
		}
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
	_ = conn.SetReadDeadline(time.Now().Add(time.Second))

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
