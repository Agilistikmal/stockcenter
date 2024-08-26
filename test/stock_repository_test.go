package test

import (
	"log"
	"testing"

	"github.com/agilistikmal/stockcenter/internal/stockcenter/infrastructure/config"
	"github.com/agilistikmal/stockcenter/internal/stockcenter/infrastructure/datastore"
	"github.com/agilistikmal/stockcenter/internal/stockcenter/repository"
)

var r *repository.StockRepository

func init() {
	config.Load()
	stockDatastore := datastore.NewStockDatastore()
	r = repository.NewStockRepository(stockDatastore)
}

func TestPush(t *testing.T) {
	err := r.Push("stock3@gmail.com:abc123")
	if err != nil {
		log.Fatal(err)
	}
}

func TestCount(t *testing.T) {
	count, err := r.Count()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Stock Count:", count)
}

func TestPopFirst(t *testing.T) {
	contentSlice, err := r.PopFirst(34)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(contentSlice)
}
