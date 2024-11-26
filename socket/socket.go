package socket

import (
	"bufio"
	"encoding/json"
	"net"
)

func Send(conn net.Conn, packet interface{}) bool {
	data, err := json.Marshal(packet)
	if err != nil {
		return false
	}

	data = append(data, '\n')
	bytes, err := conn.Write(data)
	for err == nil && bytes < len(data) {
		bytes, err = conn.Write(data[bytes:])
	}
	return bytes == len(data)
}

func Recieve[T interface{}](conn net.Conn) *T {
	reader := bufio.NewReader(conn)
	line, err := reader.ReadBytes('\n')
	if err != nil {
		return nil
	}

	output := new(T)
	err = json.Unmarshal(line, output)
	if err != nil {
		return nil
	}
	return output
}
