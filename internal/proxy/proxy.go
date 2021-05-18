package proxy

import (
	"errors"
	"fmt"
	"io"
	"net"
)

func Serve() error {
	server, err := net.Listen("tcp", ":1080")
	if err != nil {
		return err
	}

	for {
		client, err := server.Accept()
		if err != nil {
			fmt.Printf("accept failed: %v\n", err)
			continue
		}

		go process(client)
	}
}

func process(client net.Conn) {
	if err := Socks5Auth(client); err != nil {
		fmt.Printf("auth error: %v\n", err)
		client.Close()
		return
	}

	target, err := Socks5Connect(client)
	if err != nil {
		fmt.Printf("connect error: %v\n", err)
		client.Close()
		return
	}

	Socks5Forward(client, target)
}

func Socks5Auth(client net.Conn) error {
	buf := make([]byte, 256)

	// 读取 VER 和 NMETHODS
	n, err := io.ReadFull(client, buf[:2])
	if n != 2 {
		return errors.New("reading header error: " + err.Error())
	}

	ver, nMethods := int(buf[0]), int(buf[1])
	if ver != 5 {
		return errors.New(fmt.Sprintf("invalid version %d\n", ver))
	}

	// 读取 METHODS 列表
	n, err = io.ReadFull(client, buf[:nMethods])
	if n != nMethods {
		return errors.New("reading methods error: " + err.Error())
	}

	// 0x00-无需认证
	n, err = client.Write([]byte{0x05, 0x00})
	if n != 2 || err != nil {
		return errors.New("write response error: " + err.Error())
	}

	return nil
}

func Socks5Connect(client net.Conn) (net.Conn, error) {
	buf := make([]byte, 256)

	n, err := io.ReadFull(client, buf[:4])
	if n != 4 {
		return nil, errors.New("read header error: " + err.Error())
	}

	ver, cmd, _, atyp := buf[0], buf[1], buf[2], buf[3]
	if ver != 5 || cmd != 1 {
		return nil, errors.New(fmt.Sprintf("invalid ver/cmd: ver=%d cmd=%d\n", ver, cmd))
	}

	return nil, nil
}
