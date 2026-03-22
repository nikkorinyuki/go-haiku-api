FROM golang:latest AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN go build -o main ./src

# 実行ステージ
FROM gcr.io/distroless/base-debian10 AS prod

COPY --from=builder /app/main /usr/local/bin/main

CMD ["main"]