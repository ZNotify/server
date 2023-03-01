FROM golang:1.20-alpine as builder

WORKDIR /app

RUN apk --update add --no-cache ca-certificates openssl tzdata wget unzip gcc musl-dev make

COPY . .

RUN make build-production BINARY=server

FROM scratch

WORKDIR /app

ENV GIN_MODE release
ENV TZ Asia/Shanghai

EXPOSE 14444

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/bin/server /app/server

VOLUME ["/app/data/"]

ENTRYPOINT ["/app/server"]