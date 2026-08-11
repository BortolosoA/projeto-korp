FROM golang:tip-alpine3.24 AS builder

WORKDIR /app

COPY . .

RUN go build -o projeto-korp main.go

FROM scratch

COPY --from=builder /app /app

ENTRYPOINT [ "/app/projeto-korp" ]