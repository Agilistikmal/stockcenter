package service

import "github.com/agilistikmal/stockcenter/internal/stockcenter/repository"

type StockService struct {
	stockRepository *repository.StockRepository
}

func NewStockRepository(stockRepository *repository.StockRepository) *StockService {
	return &StockService{
		stockRepository: stockRepository,
	}
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
