package main

import (
	"fmt"
	"net/http"

	"desafio_taghos/internal/framework/logs"
)

func handler(w http.ResponseWriter, r *http.Request) {
	logs.Logger.Info().Msg("Recebida uma requisição HTTP")

	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	logs.InitializeLogger()

	logs.Logger.Info().Msg("Servidor rodando na porta 3000...")

	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":3000", nil); err != nil {
		logs.Logger.Fatal().Err(err).Msg("Erro ao iniciar o servidor")
	}
}
