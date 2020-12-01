package Service

import (
	"github.com/astaxie/beego/logs"
	"net"
	"sync"
	"time"
)

type UdpForward struct {
	SrcAddr          *net.UDPAddr
	DestAddr          *net.UDPAddr
	LClientAddr       *net.UDPAddr
	UdpListenerConn *net.UDPConn
	UdpConns      map[string]UdpConn
	UdpConnsMutex *sync.RWMutex
	ChkActTime time.Duration
	Closed bool
	ConnectedEvent func(addr string)
	DisConnectedEvent func(addr string)
}

type UdpConn struct {
	udp        *net.UDPConn
	lastActive time.Time
}

const bufferSize = 4096
var chkActTime = time.Minute * 1

func NewUdpForward() *UdpForward {
	return &UdpForward{
		UdpConns:make(map[string]UdpConn),
		UdpConnsMutex:new(sync.RWMutex),
		ChkActTime:chkActTime,
		ConnectedEvent:func(addr string) {},
		DisConnectedEvent:func(addr string){},
	}
}

func  (_self *UdpForward) DoUdpForward (srcAddr string, destAddr string) error {

	var err error
	_self.SrcAddr, err = net.ResolveUDPAddr("udp", srcAddr)
	if err != nil {
		logs.Error("ResolveUDPAddr ", srcAddr, " 出错：", err)
		return err
	}

	_self.DestAddr, err = net.ResolveUDPAddr("udp", destAddr)
	if err != nil {
		logs.Error("ResolveUDPAddr ", destAddr, " 出错：", err)
		return err
	}

	_self.LClientAddr = &net.UDPAddr{
		IP:   _self.SrcAddr.IP,
		Port: 0,
		Zone: _self.SrcAddr.Zone,
	}

	_self.UdpListenerConn, err = net.ListenUDP("udp", _self.SrcAddr)
	if err != nil {
		logs.Error("启动UDP监听 ", srcAddr, " 出错：", err)
		return err
	}

	go _self.checkAlive()
	go _self.runForward()

	return nil
}


func  (_self *UdpForward) runForward() {
	for {
		buf := make([]byte, bufferSize)
		n, addr, err := _self.UdpListenerConn.ReadFromUDP(buf)
		if err != nil {
			return
		}
		go _self.forwardHandler(buf[:n], addr)
	}
}

func  (_self *UdpForward) forwardHandler(data []byte, addr *net.UDPAddr) {

	_self.UdpConnsMutex.RLock()
	udpConn, found := _self.UdpConns[addr.String()]
	_self.UdpConnsMutex.RUnlock()

	if found {

		udpConn.udp.WriteTo(data, _self.DestAddr)
	}else{

		conn, err := net.ListenUDP("udp", _self.LClientAddr)
		if err != nil {
			logs.Error("udp-forwader: failed to dial:", err)
			return
		}

		_self.UdpConnsMutex.Lock()
		_self.UdpConns[addr.String()] = UdpConn{
			udp:        conn,
			lastActive: time.Now(),
		}
		_self.UdpConnsMutex.Unlock()

		_self.ConnectedEvent(addr.String())

		conn.WriteTo(data, _self.DestAddr)

		for {
			buf := make([]byte, bufferSize)
			n, _, err := conn.ReadFromUDP(buf)
			if err != nil {
				_self.UdpConnsMutex.Lock()
				conn.Close()
				delete(_self.UdpConns, addr.String())
				_self.UdpConnsMutex.Unlock()
				return
			}

			go func(data []byte, conn *net.UDPConn, addr *net.UDPAddr) {
				_self.UdpListenerConn.WriteTo(data, addr)
			}(buf[:n], conn, addr)
		}
	}

	_self.updateActiveTime(addr)
}

func  (_self *UdpForward) updateActiveTime(addr *net.UDPAddr) {

	needUpdateTime := false
	_self.UdpConnsMutex.RLock()
	if _, found := _self.UdpConns[addr.String()]; found {
		if _self.UdpConns[addr.String()].lastActive.Before(
			time.Now().Add(_self.ChkActTime / 4)) {
			needUpdateTime = true
			//logs.Debug("needUpdateTime")
		}
	}
	_self.UdpConnsMutex.RUnlock()

	if needUpdateTime {
		_self.UdpConnsMutex.Lock()
		//
		if _, found := _self.UdpConns[addr.String()]; found {
			connWrapper := _self.UdpConns[addr.String()]
			connWrapper.lastActive = time.Now()
			_self.UdpConns[addr.String()] = connWrapper
		}
		_self.UdpConnsMutex.Unlock()
	}
}

func  (_self *UdpForward) checkAlive() {

	for !_self.Closed {
		time.Sleep(_self.ChkActTime)
		var keysToDelete []string

		_self.UdpConnsMutex.RLock()
		for k, conn := range _self.UdpConns {
			if conn.lastActive.Before(time.Now().Add(-_self.ChkActTime)) {
				keysToDelete = append(keysToDelete, k)
				//logs.Debug("need delete udp conn")
			}
		}
		_self.UdpConnsMutex.RUnlock()

		_self.UdpConnsMutex.Lock()
		for _, k := range keysToDelete {
			_self.UdpConns[k].udp.Close()
			delete(_self.UdpConns, k)
		}
		_self.UdpConnsMutex.Unlock()

		for _, k := range keysToDelete {
			_self.DisConnectedEvent(k)
		}
	}

}


func  (_self *UdpForward) Close() {
	_self.UdpConnsMutex.Lock()
	_self.Closed = true
	for _, conn := range _self.UdpConns {
		conn.udp.Close()
	}
	_self.UdpListenerConn.Close()
	_self.UdpConnsMutex.Unlock()
}

func  (_self *UdpForward) GetConnsInfo() []string {
	_self.UdpConnsMutex.Lock()
	defer _self.UdpConnsMutex.Unlock()
	results := make([]string, 0, len(_self.UdpConns))
	for key := range _self.UdpConns {
		results = append(results, key)
	}
	return results
}