package service

import (
	"math"

	"github.com/agilistikmal/stockcenter/internal/stockcenter/model"
	"github.com/agilistikmal/stockcenter/internal/stockcenter/repository"
)

type StockService struct {
	stockRepository *repository.StockRepository
}

func NewStockRepository(stockRepository *repository.StockRepository) *StockService {
	return &StockService{
		stockRepository: stockRepository,
	}
}

func (s *StockService) Read(page int, limit int) (*model.StockList, error) {
	contentSlice, err := s.stockRepository.Read()
	if err != nil {
		return nil, err
	}

	page = page - 1

	start := page * limit
	if start > len(contentSlice) {
		start = len(contentSlice)
	}

	end := start + limit
	if end > len(contentSlice) {
		end = len(contentSlice)
	}

	selectedContent := contentSlice[start:end]
	pageTotal := float64(len(contentSlice)) / float64(limit)

	stockList := &model.StockList{
		Page:      page,
		Limit:     limit,
		Contents:  selectedContent,
		PageTotal: int(math.Ceil(pageTotal)),
		SizeTotal: len(contentSlice),
	}

	return stockList, nil
}

func (s *StockService) Push(content string) error {
	return s.stockRepository.Push(content)
}

func (s *StockService) PopFirst(n int) ([]string, error) {
	return s.stockRepository.PopFirst(n)
}

func (s *StockService) Replace(content string) error {
	return s.stockRepository.Replace(content)
}

func (s *StockService) Count() (int, error) {
	return s.stockRepository.Count()
}
