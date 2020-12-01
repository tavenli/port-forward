
::Build For Linux

echo "Clean for build..."
go clean

echo "Build For Linux..."

::set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64

go build -o forward-server

::go build -ldflags "-s -w" -o forward-server
::-s 去掉符号表
::-w 去掉DWARF调试信息，得到的程序不能用gdb调试

echo "--------- Build For Linux Success!"


pause

