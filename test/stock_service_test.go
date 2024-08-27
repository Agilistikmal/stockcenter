package test

import (
	"log"
	"testing"

	"github.com/agilistikmal/stockcenter/internal/stockcenter/infrastructure/config"
	"github.com/agilistikmal/stockcenter/internal/stockcenter/infrastructure/datastore"
	"github.com/agilistikmal/stockcenter/internal/stockcenter/repository"
	"github.com/agilistikmal/stockcenter/internal/stockcenter/service"
)

var s *service.StockService

func init() {
	config.Load()
	stockDatastore := datastore.NewStockDatastore()
	r := repository.NewStockRepository(stockDatastore)
	s = service.NewStockRepository(r)
}

func TestReadService(t *testing.T) {
	contentSlice, err := s.Read(4, 15)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(contentSlice)
}
