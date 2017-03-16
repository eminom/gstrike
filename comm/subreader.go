package comm

import (
	"bufio"
	"encoding/binary"
	"errors"
)

const (
	PreLength int = 4 // length includes the length of head
)

func ReceivePacket(reader *bufio.Reader) ([]byte, error) {
	p, err := reader.Peek(PreLength)
	if err != nil {
		return nil, err
	}
	totLen := binary.BigEndian.Uint32(p)
	if totLen == 0 {
		panic("Error sub read")
	}
	//log.Debugf("Incoming length:%v", totLen)
	buffer := make([]byte, totLen)
	iRd, e1 := reader.Read(buffer)
	if e1 != nil {
		return nil, e1
	}
	return buffer[PreLength:iRd], e1
}

func SendPacket(writer *bufio.Writer, buffer []byte) error {
	outBuffer := make([]byte, len(buffer)+4)
	binary.BigEndian.PutUint32(outBuffer[:4], uint32(len(buffer)+4))
	copy(outBuffer[4:], buffer)
	iw, err := writer.Write(outBuffer)
	writer.Flush() // Always remember to flush
	if err != nil {
		return err
	}
	if iw != len(outBuffer) {
		return errors.New("less than expect to write")
	}
	return nil
}
