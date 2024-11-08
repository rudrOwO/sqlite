package features

import (
	btree "github.com/rudrowo/sqlite/internal/btree"
	u "github.com/rudrowo/sqlite/internal/utils"
	"os"
)

var (
	leafPagesChannel chan btree.LeafTablePage
	dbFile           *os.File
)

func Init() *os.File {
	var err error
	dbFile, err = os.Open(os.Args[1])
	u.HandleError(err)

	leafPagesChannel = make(chan btree.LeafTablePage, 1)
	return dbFile
}
