package controller

import (
	"fmt"
	"log"
	"mongodb-com-go/database"
	"mongodb-com-go/models"
	"mongodb-com-go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BemVindo(context *gin.Context) {
	context.JSON(200, gin.H{"mensagem": "bem vindo"})
}

func CriaUsuario(context *gin.Context) {
	client :=  database.PegaClient()

	colecaoUsers := client.Database("Cluster0").Collection("users")
	var novoUsuario models.Usuario
	if err := context.ShouldBindJSON(&novoUsuario); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	senhaHasheada, err := utils.CriptografaSenha(novoUsuario.Senha)
	if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao hashear a senha"})
        return
    }
	novoUsuario.Senha = senhaHasheada
	_, err = colecaoUsers.InsertOne(context, novoUsuario)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
	}
	context.JSON(201, gin.H{"mensagem": "usuario criado com sucesso"})
}


func ListaUsuarios(context *gin.Context) {
	client := database.PegaClient()

    colecaoUsers := client.Database("Cluster0").Collection("users")

    ctx := context.Request.Context()

	cursor, err := colecaoUsers.Find(ctx, bson.M{})
	if err != nil {
		log.Panic(err.Error())
	}
	defer cursor.Close(ctx)

	var usuarios []models.Usuario
	
	for cursor.Next(ctx) {
		var usuario models.Usuario
		if err := cursor.Decode(&usuario); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
		}
		usuarios = append(usuarios, usuario)
	}
	context.JSON(http.StatusOK, usuarios)
}

func PegaUmUsuario(context *gin.Context) {
	id := context.Param("id") // Usando Param ao invés de Params.ByName
	var usuario models.Usuario
	client := database.PegaClient()
	fmt.Println(id)
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
    log.Println("Invalid id")
	}
	colecaoUsers := client.Database("Cluster0").Collection("users")
	err = colecaoUsers.FindOne(context, bson.M{"_id": objectId}).Decode(&usuario)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "usuário não encontrado", "erro": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"usuario": usuario}) // Trocado para http.StatusOK
}

func AtualizaUsuario(context *gin.Context) {
	id := context.Param("id")

	type NomeUsuario struct {
		Nome string `json:"nome"`
	}

	var novoNome NomeUsuario
	 if err := context.BindJSON(&novoNome); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := database.PegaClient()
	colecaoUsers := client.Database("Cluster0").Collection("users")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	filtro := bson.M{"_id": objectID}

	update := bson.M{"$set": bson.M{"nome": novoNome.Nome}}

	_, err = colecaoUsers.UpdateOne(context, filtro, update)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Nome do usuário atualizado com sucesso"})
}

func ApagaUsuario(context *gin.Context) {
	id := context.Param("id")
	client := database.PegaClient()
	colecaoUsers := client.Database("Cluster0").Collection("users")
	objectID, err := primitive.ObjectIDFromHex(id)
	filtro := bson.M{"_id": objectID}
	resultado, err := colecaoUsers.DeleteOne(context, filtro)
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if resultado.DeletedCount == 0 {
        context.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
        return
    }
	context.JSON(200, gin.H{"mensagem": "apagado com sucesso"})
}