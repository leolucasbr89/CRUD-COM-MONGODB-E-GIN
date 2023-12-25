package database

import (
	"context"
	"fmt"
	"log"
	"time"
    "os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConexaoComBancoDeDados() (*mongo.Client, error) {
    err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar arquivo .env:", err)
	}

	uribancodedados := os.Getenv("URLBANCODEDADOS")
    clientOpcoes := options.Client().ApplyURI(uribancodedados)

    client, err := mongo.Connect(context.Background(), clientOpcoes)
    if err != nil {
        return nil, err
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    err = client.Ping(ctx, nil)
    if err != nil {
        return nil, err
    }
    fmt.Println("Conexão com o MongoDB estabelecida")

    database := client.Database("Cluster0")
    if err := verificaExistenciaColecao(ctx, database, "users"); err != nil {
        log.Println("Erro ao verificar a existência da coleção:", err)
    }

	Client = client
    return client, nil
}

func verificaExistenciaColecao(ctx context.Context, db *mongo.Database, nomeColecao string) error {
    colecoes, err := db.ListCollectionNames(ctx, bson.M{"name": nomeColecao})
    if err != nil {
        return err
    }

    for _, nome := range colecoes {
        if nome == nomeColecao {
            fmt.Println("Coleção", nomeColecao, "já existe.")
            return nil
        }
    }

    fmt.Println("Criando coleção", nomeColecao)
    if err := db.CreateCollection(ctx, nomeColecao); err != nil {
        return err
    }

    fmt.Println("Coleção", nomeColecao, "criada com sucesso.")
    return nil
}


func PegaClient() *mongo.Client {
	return Client
}

