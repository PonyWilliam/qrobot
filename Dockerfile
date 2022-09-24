FROM golang:1.18
WORKDIR /src

COPY . /src/

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod tidy

ENTRYPOINT [ "go","run","main.go" ]

EXPOSE 9000