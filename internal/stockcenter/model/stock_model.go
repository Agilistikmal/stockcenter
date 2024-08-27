package model

type StockList struct {
	Page      int
	Limit     int
	Contents  []string
	PageTotal int
	SizeTotal int
}
