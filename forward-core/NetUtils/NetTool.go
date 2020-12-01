package NetUtils

import (
	"bufio"
	"encoding/binary"
	"errors"
	"github.com/astaxie/beego/logs"
	"io"
	"net"
	"forward-core/Common"
)

func NewTCP(addr string) (net.Listener, error) {
	tcpSocket, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	return tcpSocket,nil

}
func NewUDP(addr string) (*net.UDPConn, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}
	udpSocket, _err := net.ListenUDP("udp", udpAddr)
	if _err != nil {
		return nil, _err
	}

	return udpSocket,nil

}

func NewKCP(addr string, setting *Common.KcpSetting) (*Common.UDPListener, error) {
	return Common.NewKCP(addr, setting)
}


func DataCopy(dst io.Writer, src io.Reader) (written int64, err error) {
	return io.Copy(dst, src)
}

func MultiDataCopy(src io.Reader, dispatchConns []io.Writer) (written int64, err error) {

	mWriter := io.MultiWriter(dispatchConns...)
	return io.Copy(mWriter, src)

}

type ReadCallBack func(conn net.Conn, id int, cmd byte, arg []byte)

func ReadConn(conn net.Conn, callback ReadCallBack) {
	scanner := bufio.NewScanner(conn)
	scanner.Split(func(data []byte, atEOF bool) (adv int, token []byte, err error) {
		return NetSplitV1(data, atEOF, conn, callback)
	})
	for scanner.Scan() {
	}
	if scanner.Err() != nil {
		logs.Error(scanner.Err())
	}
}

func NetSplitV1(data []byte, atEOF bool, conn net.Conn, callback ReadCallBack) (adv int, token []byte, err error) {
	l := len(data)
	if l < 6 {
		return 0, nil, nil
	}

	if l > 100000 {
		conn.Close()
		return 0, nil, errors.New("max data")
	}
	var id int
	var cmd byte
	id = int(int32(data[0]) | int32(data[1])<<8 | int32(data[2])<<16 | int32(data[3])<<24)
	cmd = data[4]
	isShort := data[5]
	var payload []byte
	var offset int
	if isShort == 1 {
		offset = 6
	} else {
		if l < 10 {
			return 0, nil, nil
		}
		ls := binary.LittleEndian.Uint32(data[6:])
		tail := l - 10
		if tail < int(ls) {
			return 0, nil, nil
		}
		payload = data[10 : 10+ls]
		offset = 10 + int(ls)
	}

	callback(conn, id, cmd, payload)

	return offset, []byte{}, nil
}

func WriteConn(conn net.Conn, id int, cmd byte, payload []byte) error {
	if conn == nil {
		return nil
	}
	l := len(payload)
	var buf []byte
	var size int
	if l > 0 {
		size = 10 + l
		//4+1+1+4 id cmd isShort
		buf = make([]byte, size)
	} else {
		size = 6
		//4+1+1 id cmd isShort
		buf = make([]byte, size)
	}
	buf[0] = byte(id)
	buf[1] = byte(id >> 8)
	buf[2] = byte(id >> 16)
	buf[3] = byte(id >> 24)
	buf[4] = cmd
	if l > 0 {
		buf[5] = 0
		binary.LittleEndian.PutUint32(buf[6:], uint32(l))
		copy(buf[10:], []byte(payload))
	} else {
		buf[5] = 1
	}
	_, err := conn.Write(buf[:size])
	if err != nil {
		logs.Error("conn.Write error:", err)
	}
	return err
}
