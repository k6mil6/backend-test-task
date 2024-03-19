FROM golang:1.21-alpine AS builder

WORKDIR /usr/local/src

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY ./ ./
RUN go build -o ./bin/app cmd/httpgateway/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/app /app
COPY dev.env /dev.env

CMD ["/app"]