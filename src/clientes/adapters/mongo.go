package clientesadapters

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
)

const (
	// Path to the AWS CA file
	caFilePath = "src/clientes/adapters/rds-combined-ca-bundle.pem"

	// Timeout operations after N seconds
	connectTimeout  = 5
	queryTimeout    = 30
	username        = "lealadm"
	password        = "LealColombia2019%21%2B%23"
	clusterEndpoint = "leal-prod.cluster-ckbqrtxywlt9.us-west-2.docdb.amazonaws.com"

	// Which instances to read from
	readPreference = "secondaryPreferred&retryWrites=false"

	connectionStringTemplate = "mongodb://%s:%s@%s/listas?ssl=true&replicaSet=rs0&readpreference=%s"
)

var instance *mongo.Client

func GetMongoInstance() *mongo.Client {
	if instance == nil {
		startClient()
	}

	return instance
}

func startClient() {

	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint, readPreference)

	tlsConfig, err := getCustomTLSConfig(caFilePath)
	if err != nil {
		log.Fatalf("Failed getting TLS configuration: %v", err)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI).SetTLSConfig(tlsConfig))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping cluster: %v", err)
	}

	instance = client

	fmt.Println("Connected to DocumentDB!")

}

/*
func ingore() {
    collection := client.Database("listas").Collection("poc_idComercio_3")

	ctx, cancel = context.WithTimeout(context.Background(), queryTimeout*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	if err != nil {
		log.Fatalf("Failed to insert document: %v", err)
	}

	id := res.InsertedID
	log.Printf("Inserted document ID: %s", id)

	ctx, cancel = context.WithTimeout(context.Background(), queryTimeout*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		log.Fatalf("Failed to run find query: %v", err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		log.Printf("Returned: %v", result)

		if err != nil {
			log.Fatal(err)
		}
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

}

*/

func CloseMongoClient() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := instance.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func getCustomTLSConfig(caFile string) (*tls.Config, error) {
	tlsConfig := new(tls.Config)
	certs, err := ioutil.ReadFile(caFile)

	if err != nil {
		return tlsConfig, err
	}

	tlsConfig.RootCAs = x509.NewCertPool()
	ok := tlsConfig.RootCAs.AppendCertsFromPEM(certs)

	if !ok {
		return tlsConfig, errors.New("Failed parsing pem file")
	}

	return tlsConfig, nil
}
