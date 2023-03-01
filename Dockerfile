FROM --platform=$BUILDPLATFORM golang:1.20 as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM

ENV CGO_ENABLED 1
ENV GOOS linux
ENV DEBIAN_FRONTEND noninteractive

WORKDIR /app

RUN apt update && \
    apt install -y ca-certificates openssl tzdata wget unzip gcc musl-dev make && \
    update-ca-certificates && \
    if [ "$TARGETPLATFORM" = "linux/arm/v8" ]; then \
        apt install -y gcc-aarch64-linux-gnu binutils-aarch64-linux-gnu; \
    fi

COPY . .

RUN make frontend && \
    if [ "$TARGETPLATFORM" = "linux/amd64" ]; then \
        make build-production BINARY=server; \
    elif [ "$TARGETPLATFORM" = "linux/arm/v8" ]; then \
        GOARCH=arm64 CC=gcc-aarch64-linux-gnu make build-production BINARY=server; \
    fi

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