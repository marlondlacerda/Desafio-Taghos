# variáveis
# export MONGODB_URI=mongodb://localhost:27017
# export MONGO_PARAM=retryWrites=true&w=1&readPreference=primaryPreferred

export API_PORT=3000
export STAGE=development
export BATCH_SIZE=10

# Tarefas
all: build run

clean:
	rm -rf bin/

build:
	go build -tags dev -o bin/desafiotaghos cmd/main.go

run:
	./bin/session

dev:
	air

lint:
	golangci-lint run
