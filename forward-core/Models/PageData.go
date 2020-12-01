package Models

type PageData struct {
	PIndex    int64
	PSize     int64
	TotalRows int64
	Pages     int64
	Data      interface{}
}
