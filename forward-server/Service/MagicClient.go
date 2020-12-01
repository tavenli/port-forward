package Service

import "net"

type MagicClient struct {

	cid string
	encode, decode func([]byte) []byte
	authed bool
	conn net.Conn

}