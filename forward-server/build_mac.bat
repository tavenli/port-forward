

echo "Build For Mac..."
set GOOS=darwin
set GOARCH=amd64

go build -o forward-server

echo "--------- Build For Mac Success!"


pause

