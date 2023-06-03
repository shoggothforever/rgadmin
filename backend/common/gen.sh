#单体服务快速启动
goctl api go -api ./*.api -dir ../  --style=goZero
# 单体服务向微服务迁移使用
goctl rpc protoc ./*.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=goZero
#api 目录下执行
goctl model mongo --type User --dir cache --cache

cd E:\golands\rgadmin\backend\app\user\cmd\api && go run user.go -f .\etc\user.yaml
