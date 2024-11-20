FROM golang:1.23

WORKDIR /build

ENV WEBSOCKET_PORT=3002
ENV WEBSOCKET_HOST=0.0.0.0
ENV LOG_INDEX=signal-server
ENV LOG_IS_DEBUG=false
ENV LOG_TO_FILE=false

ENV GRPC_HOST=0.0.0.0
ENV GRP_PORT=50002

ENV SFU_HOST=sfu-service
ENV SFU_GRPC_PORT=50000

ENV JAEGER_HOST=jaeger:6831
ENV JAEGER_LOG_SPANS=true

COPY app /build/app
COPY cmd /build/cmd
COPY config /build/config
COPY internal /build/internal
COPY pkg /build/pkg

COPY go.mod /build/
COPY go.sum /build/


RUN go mod download && \
    go build -o cmd/signal-server/signal-server ./cmd/signal-server/main.go && \
    chmod +x cmd/signal-server/signal-server

ENTRYPOINT ["cmd/signal-server/signal-server"]

EXPOSE 3002/tcp
EXPOSE 50002/tcp