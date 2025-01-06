FROM golang:1.23-alpine

RUN apk add --no-cache ca-certificates \
    && apk add --update --upgrade curl bash make

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod tidy

WORKDIR /go/src

ENTRYPOINT ["tail", "-f", "/dev/null"]
