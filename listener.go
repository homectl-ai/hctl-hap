package hap

import (
	"net"
	"time"
)

var keepAlive = net.KeepAliveConfig{Enable: true, Idle: 30 * time.Second, Interval: 15 * time.Second, Count: 4}

type listener struct {
	*net.TCPListener
}

func (ln *listener) Accept() (con net.Conn, err error) {
	con, err = ln.AcceptTCP()
	if err != nil {
		return
	}

	// TCP keepalives: a controller that vanished (a phone asleep, gone from
	// the network) is dropped within about 90 s instead of keeping a dead
	// session open for good, which left the Home app on "No Response" until the
	// accessory restarted. Upstream turned them off.
	if tcpconn, ok := con.(*net.TCPConn); ok {
		tcpconn.SetKeepAliveConfig(keepAlive)
	}

	conn := newConn(con)
	setConn(conn.RemoteAddr().String(), conn)

	return conn, err
}

func (ln *listener) Close() error {
	return ln.TCPListener.Close()
}

func (ln *listener) Addr() net.Addr {
	return ln.TCPListener.Addr()
}
