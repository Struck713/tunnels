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
	bytes, err := conn.Write(data)
	if err != nil {
		return false
	}
	return bytes == len(data)
}

func Recieve[T interface{}](conn net.Conn) *T {
	scanner := bufio.NewScanner(conn)
	if !scanner.Scan() {
		return nil;
	}

	output := new(T)
	err := json.Unmarshal(scanner.Bytes(), output)
	if err != nil {
		return nil
	}
	return output
}
