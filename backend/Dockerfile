FROM golang:1.20
WORKDIR usr/src/app
ENV GOOS=linux
ENV CGO_ENABLED=0
ENV GO_PROXY=https://proxy.golang.com.cn,direct
ENV GO111MODULE=auto

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /rg /app/user/cmd/api/user.go

CMD ["/rg","-f","./app/user/cmd/api/etc/user.yaml"]

#FROM scratch
#
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
#COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
#ENV TZ Asia/Shanghai
#
#WORKDIR /app
#COPY --from=builder /app/user /app/user
#COPY --from=builder /app/etc /app/etc

