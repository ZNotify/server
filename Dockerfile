FROM --platform=$BUILDPLATFORM golang:1.22 as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM

ARG PREFETCHED

ENV CGO_ENABLED 1
ENV GOOS linux
ENV DEBIAN_FRONTEND noninteractive

WORKDIR /app

RUN apt update && \
    apt install -y ca-certificates openssl tzdata wget unzip gcc musl-dev make && \
    update-ca-certificates && \
    if [ "$TARGETPLATFORM" = "linux/arm64" ]; then \
        apt install -y gcc-aarch64-linux-gnu binutils-aarch64-linux-gnu; \
    fi

COPY . .

# Download frontend dependencies if not cached
RUN if [ -z "$PREFETCHED" ]; then \
        make frontend; \
    fi

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    if [ "$TARGETPLATFORM" = "linux/amd64" ]; then \
        make build-production BINARY=server; \
    elif [ "$TARGETPLATFORM" = "linux/arm64" ]; then \
        GOARCH=arm64 CC=aarch64-linux-gnu-gcc make build-production BINARY=server; \
    else \
        echo "Unsupported platform: $TARGETPLATFORM"; \
        exit 1; \
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