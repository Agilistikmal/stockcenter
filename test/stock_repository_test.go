package test

import (
	"log"
	"testing"

	"github.com/agilistikmal/stockcenter/internal/stockcenter/infrastructure/datastore"
	"github.com/agilistikmal/stockcenter/internal/stockcenter/repository"
)

var r *repository.StockRepository

func init() {
	stockDatastore := datastore.NewStockDatastore()
	r = repository.NewStockRepository(stockDatastore)
}

func TestPush(t *testing.T) {
	n, err := r.Push("stock3@gmail.com:abc123")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n)
}
