# go-haiku-api

Goで実装された、俳句判定・抽出APIです。  
`github.com/mattn/go-haiku` を使って、テキストが指定した音数ルールに一致するかを判定したり、一致する句を抽出したりできます。

## Features

- `POST /match`: テキストがルールに一致するかを `true/false` で返す
- `POST /find`: テキストからルールに一致する句を抽出して返す
- JSONレスポンス

## Requirements

- Go 1.26.1 以上

## Run (Local)

```bash
go mod download
go run main.go
```

サーバーは `:8080` で起動します。

## API

### `POST /match`

テキストがルールに一致するかを判定します。

Request:

```json
{
  "text": "古池や蛙飛び込む水の音",
  "rule": [5, 7, 5]
}
```

Response:

```json
{
  "result": true
}
```

`invalid json` の場合:

```json
{
  "error": "invalid json"
}
```

### `POST /find`

テキストからルールに一致する句を抽出します。

Request:

```json
{
  "text": "古池や蛙飛び込む水の音",
  "rule": [5, 7, 5]
}
```

Response（例）:

```json
{
  "result": [
    "古池や 蛙飛び込む 水の音"
  ]
}
```

## cURL Examples

### `/match`

```bash
curl -X POST http://localhost:8080/match \
  -H "Content-Type: application/json" \
  -d '{"text":"古池や蛙飛び込む水の音","rule":[5,7,5]}'
```

### `/find`

```bash
curl -X POST http://localhost:8080/find \
  -H "Content-Type: application/json" \
  -d '{"text":"古池や蛙飛び込む水の音","rule":[5,7,5]}'
```

## Docker

Build:

```bash
docker build -t go-haiku-api .
```

Run:

```bash
docker run --rm -p 8080:8080 go-haiku-api
```

## License

MIT