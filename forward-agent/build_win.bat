

echo "Build For windows..."
set GOOS=windows
set GOARCH=amd64

go build -o forward-agent.exe

echo "--------- Build For windows Success!"


pause

