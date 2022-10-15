FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . .

RUN apk --update add --no-cache ca-certificates openssl tzdata wget unzip gcc musl-dev

RUN wget https://github.com/ZNotify/frontend/releases/download/bundle/build.zip && \
          unzip build.zip && \
          rm build.zip && \
          mv build web/static

RUN go build -v -o /app/server -ldflags "-s -w" notify-api

FROM scratch

WORKDIR /app

ENV GIN_MODE release
ENV TZ Asia/Shanghai

EXPOSE 14444

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/server ./server

VOLUME ["/app/data/"]

ENTRYPOINT ["./server"]