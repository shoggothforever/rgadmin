goctl api go -api *.api -dir ../  --style=goZero
goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=goZero