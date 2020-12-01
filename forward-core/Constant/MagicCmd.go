package Constant

const (
	MagicCmd_AuthFail =byte(iota)
	MagicCmd_ReqAuth
	MagicCmd_Refused
	MagicCmd_AgentListenerOpen
	MagicCmd_AgentConnOpen
	MagicCmd_AgentConnClose
	MagicCmd_DataToMagic
	MagicCmd_DataToAgent
)