package dialer

import (
	"crypto/tls"
	"errors"
	"net"
	"syscall"
	"time"
)

func NewConn(addr string, timeout time.Duration) (conn net.Conn, err error) {
	conn, err = net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return
	}
	return
}

func NewTLSConn(conn net.Conn) (tlsConn *tls.Conn, err error) {
	config := &tls.Config{
		InsecureSkipVerify: true,
		MinVersion:         tls.VersionTLS10,
		MaxVersion:         tls.VersionTLS13,
	}
	tlsConn = tls.Client(conn, config)
	err = tlsConn.Handshake()
	return
}

func Send(conn net.Conn, data []byte, timeout time.Duration) error {
	err := conn.SetWriteDeadline(time.Now().Add(timeout))
	if err != nil {
		return err
	}
	_, err = conn.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func Recv(conn net.Conn, timeout time.Duration) ([]byte, error) {
	response := make([]byte, 4096)
	err := conn.SetReadDeadline(time.Now().Add(timeout))
	if err != nil {
		return []byte{}, err
	}
	length, err := conn.Read(response)
	if err != nil {
		var netErr net.Error
		if (errors.As(err, &netErr) && netErr.Timeout()) ||
			errors.Is(err, syscall.ECONNREFUSED) { // timeout error or connection refused
			return []byte{}, nil
		}
		return response[:length], err
	}
	return response[:length], nil
}
