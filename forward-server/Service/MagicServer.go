package Service

import "forward-core/Models"

type MagicServer struct {
	UseUDP bool
}

func (_self *MagicServer) StartMagicService(netAddr string, result chan Models.FuncResult) {

}