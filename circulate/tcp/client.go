package tcp

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
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
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")

	send(conn)
	recv(conn)
}
