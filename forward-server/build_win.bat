
::set GoDevWork="D:\CodeWork\port-forward\"

echo "Clean for build..."
go clean

echo "Build For windows..."

set GOOS=windows
set GOARCH=amd64
::set GOPATH=%GoDevWork%;%GOPATH%
go build -o forward-server.exe

echo "--------- Build For windows Success!"


pause

