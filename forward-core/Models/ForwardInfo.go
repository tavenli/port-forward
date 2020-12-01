package Models

type ForwardInfo struct {

	Name string
	Status byte
	SrcAddr string
	SrcPort int
	Protocol string
	DestAddr string
	DestPort int
	OnlineCount int
	Clients []string

}