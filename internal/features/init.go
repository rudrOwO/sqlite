package features

import (
	"log"
	"os"

	btree "github.com/rudrowo/sqlite/internal/btree"
)

var (
	leafPagesChannel chan btree.LeafTablePage
	dbFile           *os.File
)

func Init() *os.File {
	var err error
	dbFile, err = os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	leafPagesChannel = make(chan btree.LeafTablePage, 1)
	return dbFile
}
