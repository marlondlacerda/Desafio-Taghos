# Tarefas
all: build run

clean:
	rm -rf bin/

build:
	go build -tags dev -o bin/desafiotaghos cmd/main.go

run:
	./bin/desafiotaghos

dev:
	air

swagger:
	swag init -g cmd/main.go --parseDependency --parseInternal

lint:
	golangci-lint run
