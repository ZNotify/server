FROM golang:latest as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

ENV DEBIAN_FRONTEND=noninteractive

WORKDIR /app

COPY . .

RUN apt update && apt install -y ca-certificates openssl git tzdata wget unzip && \
update-ca-certificates

RUN wget https://github.com/ZNotify/frontend/releases/download/bundle/build.zip && \
          unzip build.zip && \
          rm build.zip && \
          mv build web/static

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