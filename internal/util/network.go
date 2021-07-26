package util

import (
	"net"
	"strings"
)

func GetOutBoundIP() (string, error) {

	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return "", err
	}

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip := strings.Split(localAddr.String(), ":")[0]

	return ip, nil
}
