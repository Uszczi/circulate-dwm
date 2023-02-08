package server

import (
	"bytes"
	"circulate/usecase"
	"circulate/win"
	"encoding/gob"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

type Message struct {
	Data string
}

func send(conn net.Conn, message string) {
	msg := Message{Data: message}
	bin_buf := new(bytes.Buffer)

	gobobj := gob.NewEncoder(bin_buf)
	_ = gobobj.Encode(msg)

	_, _ = conn.Write(bin_buf.Bytes())
}

func recv(conn net.Conn) *Message {
	tmp := make([]byte, 500)
	_, _ = conn.Read(tmp)

	tmpbuff := bytes.NewBuffer(tmp)
	tmpstruct := new(Message)

	gobobjdec := gob.NewDecoder(tmpbuff)
	_ = gobobjdec.Decode(tmpstruct)
	return tmpstruct

}

func SendCommand(message ...string) {
	log.Printf("tcp.SendCommand message=%v", message)

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

func read(conn net.Conn) {
	var msg *Message

	for {
		msg = recv(conn)
		args := strings.Split(msg.Data, " ")
		log.Printf("[TCP] Got message=%v", msg.Data)

		switch args[0] {
		case "debug-workspace":
			usecase.PrintWorkspaceDebug()
		case "clear-workspace":
			usecase.ClearWorkspace()
		case "switch-to-workspace":
			workspace, _ := strconv.Atoi(args[1])
			usecase.SwitchToWorkspace(workspace - 1)
		case "move-to-workspace":
			workspace, _ := strconv.Atoi(args[1])
			foregroundWindow := win.GetForegroundWindow()
			usecase.MoveToWorkspace(foregroundWindow, workspace-1)
		case "set-layout":
			switch args[1] {
			case "next":
				usecase.SetNextLayout()
			case "previous":
				usecase.SetPreviousLayout()
			default:
				usecase.SetLayout(args[1])
			}
		default:
			log.Printf("Unknow command args=%v\n", args)
		}

		return
	}
}

func resp(conn net.Conn) {
	msg := Message{Data: "Hello back"}
	bin_buf := new(bytes.Buffer)

	gobobje := gob.NewEncoder(bin_buf)
	_ = gobobje.Encode(msg)

	_, _ = conn.Write(bin_buf.Bytes())
	conn.Close()
}

func handle(conn net.Conn) {
	_ = conn.SetReadDeadline(time.Now().Add(time.Second))

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
