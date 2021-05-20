package clientesadapters

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"

	clientes "leal.co/clientes-comercio/src/clientes/domain"
)

type ClientesRepository interface {
	SaveCliente(cliente clientes.Cliente) error
}

type clientesRepository struct {
	client *mongo.Client
}

func NewClientesRepository() ClientesRepository {
	return &clientesRepository{
		client: GetMongoInstance(),
	}
}

func (c *clientesRepository) SaveCliente(cliente clientes.Cliente) error {
	bsonCliente, err := bson.Marshal(cliente)
	if err != nil {
		return err
	}
	collection := c.client.Database("listas").Collection("poc_idComercio_3")

	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{{"uid": cliente.Uid}}
	update := bson.D{{"$set", bsonCliente}}

	res, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	if res.UpsertedCount != 0 {
		log.Printf("inserted a new document with ID %v\n", res.UpsertedID)
	}

	log.Println("inserted document with id: ", res.UpsertedID)
	return nil
}
