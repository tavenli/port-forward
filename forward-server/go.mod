module forward-server

go 1.14

require (
	forward-core v0.0.0-00010101000000-000000000000
	github.com/astaxie/beego v1.12.2
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/vmihailenco/msgpack v4.0.4+incompatible
	google.golang.org/appengine v1.6.7 // indirect
)

replace forward-core => ../forward-core
