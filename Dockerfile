FROM golang:1.23-alpine

RUN apk add --no-cache ca-certificates \
    && apk add --update --upgrade curl bash make

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go install github.com/air-verse/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/main.go --parseDependency --parseInternal

WORKDIR /go/src

ENTRYPOINT ["tail", "-f", "/dev/null"]
