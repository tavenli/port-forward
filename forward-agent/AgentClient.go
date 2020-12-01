package main

import "net"

type AgentClient struct {

	cid string
	encode, decode func([]byte) []byte
	authed bool
	conn net.Conn

}