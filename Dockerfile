FROM debian:13 AS builder

WORKDIR /app

COPY go.* ./

RUN apt-get update && apt-get upgrade -y && apt-get install -y curl gcc
RUN ARCH=$(dpkg --print-architecture) && \
if [ "$ARCH" = "arm64" ]; then GOARCH="arm64"; else GOARCH="amd64"; fi && \
curl -OL https://go.dev/dl/go1.26.1.linux-${GOARCH}.tar.gz && \
tar -C /usr/local -xzf go1.26.1.linux-${GOARCH}.tar.gz && \
rm -rf go1.26.1.linux-${GOARCH}.tar.gz
ENV PATH $PATH:/usr/local/go/bin
ADD ./ .

RUN go mod download
RUN go build -o main main.go

# 実行ステージ
FROM gcr.io/distroless/base-debian13 AS prod

COPY --from=builder /app/main /usr/local/bin/main

CMD ["main"]
