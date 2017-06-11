package common

import (
	"bufio"
	"encoding/binary"
	"errors"
	"net"

	"github.com/astaxie/beego/logs"
)

type NetCommon struct {
}

type ReadCallBack func(conn net.Conn, id int, cmd byte, arg []byte)

const (
	MagicAuthFail = byte(iota)
	DoAuth
	MagicRefused
	AgentListenerOpen
	AgentConnOpen
	AgentConnClose
	MsgToMagic
	MsgToAgent
)

func (_self *NetCommon) ReadConn(conn net.Conn, callback ReadCallBack) {
	scanner := bufio.NewScanner(conn)
	scanner.Split(func(data []byte, atEOF bool) (adv int, token []byte, err error) {
		return _self.Split(data, atEOF, conn, callback)
	})
	for scanner.Scan() {
	}
	if scanner.Err() != nil {
		logs.Error(scanner.Err())
	}
}

func (_self *NetCommon) Split(data []byte, atEOF bool, conn net.Conn, callback ReadCallBack) (adv int, token []byte, err error) {
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

func (_self *NetCommon) WriteConn(conn net.Conn, id int, cmd byte, payload []byte) error {
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
