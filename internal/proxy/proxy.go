package proxy

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

var verbose1 bool

type ProxyServer struct {
	port    uint
	verbose bool
}

func NewProxyServer(port uint, verbose bool) *ProxyServer {
	verbose1 = verbose
	return &ProxyServer{port, verbose}
}

func (p *ProxyServer) Serve() error {
	server, err := net.Listen("tcp", fmt.Sprintf(":%d", p.port))
	if err != nil {
		return err
	}

	fmt.Printf("Proxy Server Running on %d\n", p.port)

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
		return fmt.Errorf("invalid version %d", ver)
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

	// 在完成认证以后，客户端需要告知服务端它的目标地址，协议具体要求为：

	// --------------------------------------------------------
	// VER	CMD	RSV		ATYP	DST.ADDR	DST.PORT
	// 1	1	X'00'	1		Variable	2
	// --------------------------------------------------------
	// VER 		0x05，老暗号了
	// CMD 		连接方式，0x01=CONNECT, 0x02=BIND, 0x03=UDP ASSOCIATE
	// RSV 		保留字段，现在没卵用
	// ATYP 	地址类型，0x01=IPv4，0x03=域名，0x04=IPv6
	// DST.ADDR	目标地址，细节后面讲
	// DST.PORT	目标端口，2字节，网络字节序（network octec order）
	buf := make([]byte, 256)

	n, err := io.ReadFull(client, buf[:4])
	if n != 4 {
		return nil, errors.New("read header error: " + err.Error())
	}

	ver, cmd, _, atyp := buf[0], buf[1], buf[2], buf[3]
	if ver != 5 || cmd != 1 {
		return nil, fmt.Errorf("invalid ver/cmd: ver=%d cmd=%d", ver, cmd)
	}

	// 接下来问题是如何读取 DST.ADDR 和 DST.PORT。

	// 如前所述，ADDR 的格式取决于 ATYP：
	// 0x01：4个字节，对应 IPv4 地址
	// 0x03：先来一个字节 n 表示域名长度，然后跟着 n 个字节。注意这里不是 NUL 结尾的。
	// 0x04：16个字节，对应 IPv6 地址

	addr := ""
	switch atyp {
	case 1:
		n, err := io.ReadFull(client, buf[:4])
		if n != 4 {
			return nil, errors.New("invalid IPv4: " + err.Error())
		}
		addr = fmt.Sprintf("%d.%d.%d.%d", buf[0], buf[1], buf[2], buf[3])
	case 3:
		n, err = io.ReadFull(client, buf[:1])
		if n != 1 {
			return nil, errors.New("invalid hostname: " + err.Error())
		}
		addrLen := int(buf[0])

		n, err = io.ReadFull(client, buf[:addrLen])
		if n != addrLen {
			return nil, errors.New("invalid hostname: " + err.Error())
		}
		addr = string(buf[:addrLen])
	case 4:
		return nil, errors.New("IPv6: no supported yet")
	default:
		return nil, errors.New("invalid atyp")
	}

	// 接着要读取的 PORT 是一个 2 字节的无符号整数。
	n, err = io.ReadFull(client, buf[:2])
	if n != 2 {
		return nil, errors.New("read port: " + err.Error())
	}
	port := binary.BigEndian.Uint16(buf[:2])

	destAddrPort := fmt.Sprintf("%s:%d", addr, port)
	dest, err := net.Dial("tcp", destAddrPort)
	if err != nil {
		return nil, fmt.Errorf("dial dst[%s] error: %v", destAddrPort, err)
	}

	// 最后一步是告诉客户端，我们已经准备好了
	// --------------------------------------------------
	// VER	REP	RSV		ATYP	BND.ADDR	BND.PORT
	// 1	1	X'00'	1		Variable	2
	// --------------------------------------------------
	// VER 		暗号，还是暗号！
	// REP 		状态码，0x00=成功，0x01=未知错误，……
	// RSV 		依然是没卵用的 RESERVED
	// ATYP 	地址类型
	// BND.ADDR 服务器和DST创建连接用的地址
	// BND.PORT 服务器和DST创建连接用的端口
	_, err = client.Write([]byte{0x05, 0x00, 0x00, 0x01, 0, 0, 0, 0, 0, 0})
	if err != nil {
		return nil, errors.New("write resp error: " + err.Error())
	}

	return dest, nil
}

func Socks5Forward(client, target net.Conn) {
	forward := func(src, dest net.Conn) {
		defer src.Close()
		defer dest.Close()
		if verbose1 {
			rs := io.TeeReader(src, dest)
			now := time.Now().Format("15:04:05")
			fmt.Printf("------------------- %s -------------------\n", now)
			io.Copy(os.Stdout, rs)
		} else {
			buf := make([]byte, 2048)
			io.CopyBuffer(src, dest, buf)
		}
	}

	go forward(client, target)
	go forward(target, client)
}
