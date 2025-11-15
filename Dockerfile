FROM golang:1.25-alpine AS build

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o qa_service ./cmd/app

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

FROM alpine:3.18

RUN apk add --no-cache bash libc6-compat

WORKDIR /app

COPY --from=build /app/qa_service .
COPY --from=build /app/migrations ./migrations
COPY --from=build /go/bin/goose /usr/local/bin/goose
COPY --from=build /app/.env .env

EXPOSE 8080

CMD ["./qa_service"]
