@echo off

cd protobuf
protoc --go_out=. --go-grpc_out=./ *.proto
echo "successed..."

pause