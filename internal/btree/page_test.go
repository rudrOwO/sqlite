package btree

import (
	"os"
	"testing"
)

func TestLoadPage(t *testing.T) {
	dbFile, err := os.Open("../../superheroes.db")
	if err != nil {
		t.Errorf(`Error Opening db file`)
	}
	defer dbFile.Close()

	fileBuffer := make([]byte, PAGE_SIZE)
	_, err = dbFile.Seek(4096, 0)
	if err != nil {
		t.Errorf(`Error Reading File`)
	}
	_, err = dbFile.Read(fileBuffer)
	if err != nil {
		t.Errorf(`Error Reading File`)
	}

	l := new(interiorTablePage)
	l.loadFromBuffer(fileBuffer)
	t.Logf("\n%+v\n", *l)
}
