package main

import (
	"fmt"
	"os"

	"github.com/rudrowo/sqlite/internal/features"
)

// Usage: your_sqlite3.sh sample.db .dbinfo
func main() {
	dbFile := features.Init()
	defer dbFile.Close()
	userCommand := os.Args[2]

	// TODO  Remove switch case and call HandleDotCommands() and HandleSelectQuery()
	switch userCommand {
	case ".dbinfo":
		fmt.Printf("database page size: %v\n", features.ReadPageSize())
		fmt.Printf("number of tables: %v", features.CountRows("sqlite_schema"))
	default:
		fmt.Println("Unknown command", userCommand)
		os.Exit(1)
	}
}
