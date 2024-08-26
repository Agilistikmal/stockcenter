package datastore

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

type StockDatastore struct {
	filepath string
}

func NewStockDatastore() *StockDatastore {
	// Data Folder
	// Create folder "data" if not exists
	dirpath := filepath.Join("./data")
	err := os.MkdirAll(dirpath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// Stocks File
	// Create "stocks.txt" file inside folder "data" if not exists
	fpath := filepath.Join("./data/stocks.txt")
	_, err = os.Stat(fpath)
	if errors.Is(err, os.ErrNotExist) {
		log.Println("Creating new stock file ./data/stocks.txt")
		err := os.WriteFile(fpath, []byte{}, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &StockDatastore{
		filepath: fpath,
	}
}

func (d StockDatastore) Open() (*os.File, error) {
	file, err := os.OpenFile(d.filepath, os.O_APPEND|os.O_RDWR, os.ModeAppend)
	return file, err
}
