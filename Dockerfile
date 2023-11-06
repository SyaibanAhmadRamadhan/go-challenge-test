FROM golang:1.21 AS build
ENV CGO_ENABLED=0

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest AS final

WORKDIR /app

COPY --from=build /app/main .
COPY --from=build /app/.env .

CMD ["sh", "-c", "./main migrate up && ./main web"]
