FROM golang:1.15.7 AS build

RUN go env -w GOPROXY=https://goproxy.cn,direct
ENV BUILD_TAGS=jsoniter \
    BUILD_LDFLAGS="-s -w -linkmode external -extldflags \"-static\""

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -tags="$BUILD_TAGS" -ldflags="$BUILD_LDFLAGS" -o ib-auth cmd/main.go

FROM alpine:3.13.1

ENV TZ=Asia/Shanghai \
    LANG=C.UTF-8

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add -u ca-certificates tzdata && \
    cp /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone

WORKDIR /app

COPY --from=build /app/ib-auth .

CMD ["./ib-auth"]