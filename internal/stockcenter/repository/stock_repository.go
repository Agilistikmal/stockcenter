package repository

import "github.com/agilistikmal/stockcenter/internal/stockcenter/infrastructure/datastore"

type StockRepository struct {
	stockDatastore *datastore.StockDatastore
}

func NewStockRepository(stockDatastore *datastore.StockDatastore) *StockRepository {
	return &StockRepository{
		stockDatastore: stockDatastore,
	}
}

func (s *StockRepository) Push(content string) (int, error) {
	file, err := s.stockDatastore.Open()
	if err != nil {
		return 0, err
	}
	defer file.Close()

	n, err := file.WriteString("\n" + content)
	return n, err
}
