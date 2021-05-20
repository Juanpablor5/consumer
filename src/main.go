package main

import (
	"log"

	"github.com/joho/godotenv"
	clientesadapters "leal.co/clientes-comercio/src/clientes/adapters"
	clientesports "leal.co/clientes-comercio/src/clientes/ports"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		log.Printf("error loading .env, %v", err)
	}
}

func main() {
	clientesPorts := clientesports.NewKinesisConsumer()
	clientesPorts.StartConsumer()
	defer clientesadapters.CloseMongoClient()

}
