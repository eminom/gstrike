package nucleo

import (
	"bufio"
	"github.com/eminom/gstrike/comm"
	"net"
)

type XClient struct {
	conn net.Conn
	wr   *bufio.Writer
	rd   *bufio.Reader
}

func NewClient(conn net.Conn) *XClient {
	return &XClient{
		conn,
		bufio.NewWriter(conn),
		bufio.NewReader(conn),
	}
}

func (xc *XClient) StartServe() {
	go func() {
		//log.Tracef("Starting serve one client.")
		rd, wr := xc.rd, xc.wr
	A100:
		for {
			buffer, err := comm.ReceivePacket(rd)
			if err != nil {
				break A100
			}
			err = comm.SendPacket(wr, buffer)
			if err != nil {
				break A100
			}
		}
		xc.conn.Close()
		//log.Tracef("Client shutdown.")
	}()
}
