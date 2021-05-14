package main

import (
	"context"
	"encoding/json"
	"expvar"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/joho/godotenv"
	consumer "github.com/mitooos/kinesis-consumer"
	storage "github.com/mitooos/kinesis-consumer/store/ddb"
)

// A myLogger provides a minimalistic logger satisfying the Logger interface.
type myLogger struct {
	logger *log.Logger
}

// Log logs the parameters to the stdlib logger. See log.Println.
func (l *myLogger) Log(args ...interface{}) {
	l.logger.Println(args...)
}

type Restaurant struct {
	Name      string `json:"nombre"`
	Direccion string `json:"direccion"`
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		log.Printf("error loading .env, %v", err)
	}
}

func main() {
	// Wrap myLogger around  apex logger
	logger := &myLogger{
		logger: log.New(os.Stdout, "consumer-example: ", log.LstdFlags),
	}

	stream := os.Getenv("STREAM")
	app := os.Getenv("APP_NAME")
	table := os.Getenv("CHECKPOINT_TABLE")

	// New Kinesis and DynamoDB clients (if you need custom config)
	// client

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-west-2"),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: os.Getenv("AWS_ACCESS_ID"), SecretAccessKey: os.Getenv("AWS_ACCESS_KEY"),
				Source: "environment variables credentials",
			},
		}),
	)
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	myDdbClient := dynamodb.NewFromConfig(cfg)

	myKsis := kinesis.NewFromConfig(cfg)

	// ddb persitance
	ddb, err := storage.New(app, table, storage.WithDynamoClient(myDdbClient))
	if err != nil {
		logger.Log("checkpoint error: ", err)
	}

	// expvar counter
	var counter = expvar.NewMap("counters")

	// consumer
	c, err := consumer.New(
		stream,
		consumer.WithStore(ddb),
		consumer.WithLogger(logger),
		consumer.WithCounter(counter),
		consumer.WithClient(myKsis),
		consumer.WithAggregation(true),
	)
	if err != nil {
		logger.Log("consumer error: ", err)
	}

	// use cancel func to signal shutdown
	ctx, cancel := context.WithCancel(context.Background())

	// trap SIGINT, wait to trigger shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals
		cancel()
	}()

	// scan stream
	err = c.Scan(ctx, func(r *consumer.Record) error {
		var responseStruct struct {
			Event string `json:"event"`
		}

		if err := json.Unmarshal(r.Data, &responseStruct); err != nil {
			return fmt.Errorf("error parsing json event, %v", err)
		}

		switch responseStruct.Event {
		case "cliente_actualizado":
			log.Println("cliente_actualizado", string(r.Data))
		case "puntos_cargados":
			log.Println("puntos_cargados", string(r.Data))
		case "usuario_comercio_actualizado":
			log.Println("usuario_comercio_actualizado", string(r.Data))
		}
		return nil

	})

	if err != nil {
		logger.Log("scan error: ", err)
	}

	if err := ddb.Shutdown(); err != nil {
		logger.Log("storage shutdown error: ", err)
	}
}
