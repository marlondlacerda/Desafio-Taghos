package database

import (
	"context"
	"desafio_taghos/infra/config"
	"desafio_taghos/internal/framework/logs"
	"encoding/json"
	"errors"
	"log/slog"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConfig define as configurações necessárias para conexão com MongoDB
type MongoConfig struct {
	URL       string      // URL de conexão com o MongoDB
	AppName   string      // Nome da aplicação
	DebugMode bool        // Flag que habilita os logs de debug
	Log       slog.Logger // Logger que será utilizado
}

// Variável global que mantém a referência do banco de dados MongoDB
var (
	MongoDB *mongo.Database
)

// Inicializa a conexão com MongoDB
//
// Se isTest for true, cria um container de teste, caso contrário usa configuração de produção
func NewMongoDBDatabaseConnection(isTest bool) error {
	if isTest {
		return errors.New("test is not defined")
	}
	return setupProductionMongoDB()
}

// Configura e estabelece conexão com MongoDB de produção
func setupProductionMongoDB() error {
	// Inicializa configurações a partir das variáveis de ambiente
	cfg := MongoConfig{}
	cfg.URL = config.Env("MONGOURI")
	cfg.AppName = config.ServerName
	cfg.DebugMode = config.Env("ENVIRONMENT") == "dev"
	cfg.Log = *slog.Default()

	// Configura opções de conexão
	options := options.Client().ApplyURI(cfg.URL)
	options.SetAppName(cfg.AppName)

	// Se modo debug está ativo, configura monitor para logging de comandos
	if cfg.DebugMode {
		monitor := &event.CommandMonitor{
			// Monitor para comandos iniciados
			Started: func(_ context.Context, e *event.CommandStartedEvent) {
				// Ignora comandos de endSessions e ping
				if e.CommandName != "endSessions" && e.CommandName != "ping" {
					// Formata e loga o comando em JSON
					command := e.Command.String()
					var commandJson map[string]interface{}
					err := json.Unmarshal([]byte(command), &commandJson)
					if err != nil {
						cfg.Log.Error("Error unmarshalling command", "error", err)
						return
					}
					r, _ := json.MarshalIndent(&commandJson, "", "  ")
					cfg.Log.Info(string(r))
				}
			},
			// Monitor para comandos bem-sucedidos
			Succeeded: func(_ context.Context, e *event.CommandSucceededEvent) {
				// Similar ao Started, mas para respostas bem-sucedidas
				if e.CommandName != "endSessions" && e.CommandName != "ping" {
					command := e.Reply.String()
					var commandJson map[string]interface{}
					err := json.Unmarshal([]byte(command), &commandJson)
					if err != nil {
						cfg.Log.Error("Error unmarshalling command", "error", err)
						return
					}
					r, _ := json.MarshalIndent(&commandJson, "", "  ")
					cfg.Log.Info(string(r))
				}
			},
			Failed: func(context.Context, *event.CommandFailedEvent) {},
		}
		options.SetMonitor(monitor)
	}

	logs.Logger.Info().Msg(options.GetURI())

	// Estabelece conexão com o MongoDB
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		return err
	}

	// Seleciona o banco de dados específico
	db := client.Database(config.Env("MONGOURI_DATABASE"))

	// Testa a conexão com ping
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	// Atribui a conexão à variável global
	MongoDB = db
	return nil
}
