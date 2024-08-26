package repository

import (
	"fmt"
	"strings"

	"github.com/agilistikmal/stockcenter/internal/stockcenter/infrastructure/datastore"
	"github.com/spf13/viper"
)

type StockRepository struct {
	stockDatastore *datastore.StockDatastore
}

func NewStockRepository(stockDatastore *datastore.StockDatastore) *StockRepository {
	return &StockRepository{
		stockDatastore: stockDatastore,
	}
}

func (s *StockRepository) Push(content string) error {
	file, err := s.stockDatastore.OpenAppend()
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("\n" + content)
	return err
}

func (s *StockRepository) Replace(content string) error {
	file, err := s.stockDatastore.OpenTruncate()
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

func (s *StockRepository) Count() (int, error) {
	content, err := s.stockDatastore.Read()
	if err != nil {
		return 0, err
	}

	separator := viper.GetString("stock.separator")
	contentSlice := strings.Split(content, separator)

	return len(contentSlice), nil
}

func (s *StockRepository) PopFirst(n int) ([]string, error) {
	content, err := s.stockDatastore.Read()
	if err != nil {
		return nil, err
	}

	separator := viper.GetString("stock.separator")
	contentSlice := strings.Split(content, separator)

	if len(contentSlice) == 1 {
		if contentSlice[0] == "" {
			return nil, fmt.Errorf("insufficient stock, available %d stocks", 0)
		}
	}

	if len(contentSlice) < n {
		return nil, fmt.Errorf("insufficient stock, available %d stocks", len(contentSlice))
	}

	selectedContent := contentSlice[:n]
	unselectedContent := contentSlice[n:]

	err = s.Replace(strings.Join(unselectedContent, separator))
	if err != nil {
		return nil, err
	}

	return selectedContent, nil
}
