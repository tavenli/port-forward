package Models

type ForwardConfig struct {
	RuleId   int
	Name     string
	Status   int
	SrcAddr  string
	SrcPort  int
	Protocol string
	DestAddr string
	DestPort int
	Others   string
}
