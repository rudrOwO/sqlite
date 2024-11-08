package features

import (
	"encoding/binary"

	u "github.com/rudrowo/sqlite/internal/utils"
)

func ReadPageSize() uint16 {
	dbHeader := make([]byte, 100)
	_, err := dbFile.Read(dbHeader)
	u.HandleError(err)

	pageSize := binary.BigEndian.Uint16(dbHeader[16:18])
	return pageSize
}

// TODO  Parse sqlite_schema and read table names
func ReadTableNames() []string {
	return []string{}
}
