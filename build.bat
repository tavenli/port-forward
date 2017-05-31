
set GoDevWork="E:\Taven\MyHobbyWork\PortForward\"

echo "Build For Linux..."
set GOOS=linux
set GOARCH=amd64
set GOPATH=%GoDevWork%;%GOPATH%
go build -o PortForward

echo "--------- Build For Linux Success!"

echo "Build For Win..."
:: set GOOS=windows
:: set GOARCH=386
:: go build -o PortForward.exe

set GOOS=windows
set GOARCH=amd64
set GOPATH=%GoDevWork%;%GOPATH%
go build -o PortForward.exe

echo "--------- Build For Win Success!"

pause

