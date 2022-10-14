FROM golang:1.19-alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

ENV DEBIAN_FRONTEND=noninteractive

WORKDIR /app

COPY . .

RUN apk --update add --no-cache ca-certificates openssl tzdata && \
    update-ca-certificates

RUN go build -v -o /app/server

FROM scratch

WORKDIR /app

ENV GIN_MODE release

EXPOSE 14444

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/server ./server

VOLUME ["/app/data/"]

ENTRYPOINT ["./server"]