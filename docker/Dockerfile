FROM golang:1.20.0-alpine3.16 as builder
RUN apk update && \
    apk add ca-certificates
ENV TZ=Europe/Warsaw
RUN apk add tzdata && \
    cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
WORKDIR /build
ENV GO111MODULE=on
ENV CGO_ENABLED=0
# ENV GOPROXY=https://goproxy.cn
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags "-w -s"

FROM alpine:3.18.0
RUN apk update && apk upgrade --no-cache
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENV TZ=Europe/Warsaw
COPY --from=builder /etc/localtime /etc/localtime
COPY --from=builder /etc/timezone /etc/timezone
COPY --from=builder /build/apt-proxy /usr/bin/apt-proxy
EXPOSE 3142/tcp
ENTRYPOINT ["/usr/bin/apt-proxy"]
