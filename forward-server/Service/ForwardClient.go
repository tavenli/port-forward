package Service

import (
	"io"
	"net"

	"github.com/astaxie/beego/logs"
)

type ForWardClient struct {
	SrcConn        net.Conn
	DestConn       net.Conn
	DispatchConns  []io.Writer
	ClosedCallBack func(srcConn net.Conn, destConn net.Conn)
}

func (_self *ForWardClient) StartForward() {

	//io.Copy(dst, src)

	go func() {
		_, err := io.Copy(_self.DestConn, _self.SrcConn)
		if err != nil {
			//logs.Error("客户端来源数据转发到目标端口异常：", err)
			_self.StopForward()
		}
	}()

	go func() {
		_, err := io.Copy(_self.SrcConn, _self.DestConn)
		if err != nil {
			//logs.Error("目标端口返回响应数据异常：", err)
			_self.StopForward()
		}
	}()
}

func (_self *ForWardClient) DispatchData(dispatchConns []io.Writer) {
	//将数据克隆给其它端口
	go func() {
		mWriter := io.MultiWriter(append(dispatchConns, _self.DestConn)...)
		_, err := io.Copy(mWriter, _self.SrcConn)
		if err != nil {
			logs.Error("Dispatch网络连接异常：", err)
			_self.StopForward()
		}
	}()

	go func() {
		_, err := io.Copy(_self.SrcConn, _self.DestConn)
		if err != nil {
			//logs.Error("目标端口返回响应数据异常：", err)
			_self.StopForward()
		}
	}()
}

func (_self *ForWardClient) StopForward() {
	logs.Debug("关闭一个连接：", _self.SrcConn.RemoteAddr(), " on ", _self.SrcConn.LocalAddr())
	_self.SrcConn.Close()
	_self.DestConn.Close()
	_self.ClosedCallBack(_self.SrcConn, _self.DestConn)

}
