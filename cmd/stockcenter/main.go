package main

import (
	"log"

	"github.com/agilistikmal/stockcenter/internal/stockcenter/infrastructure/datastore"
)

func main() {
	datastore.NewStockDatastore()
	log.Println("Stockcenter")
}
