package btree

import (
	"os"
	"testing"
)

func TestGetRootPageOffset(t *testing.T) {
	schemaRootOffset := getRootPageOffset("sqlite_schema")

	if schemaRootOffset != 100 {
		t.Errorf(`Test Failed for TestGetRootPageOffset
	offset found: %d
	`, schemaRootOffset)
	}
}

func TestLoadAllLeafTablePages(t *testing.T) {
	dbFile, err := os.Open("../../superheroes.db")
	if err != nil {
		t.Errorf(`Error Opening db file`)
	}
	defer dbFile.Close()

	testChannel := make(chan LeafTablePage, 1)
	go LoadAllLeafTablePages("sqlite_schema", dbFile, testChannel)

	count := uint16(0)
	for c := range testChannel {
		count += c.Header.CellsCount
	}

	t.Logf("\n%+v\n", count)
}
