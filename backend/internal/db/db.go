package db

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
)

var initalized bool

var path string

//go:embed schema/*
var schema embed.FS

//go:embed generic_bookmarks.json
var bookmarks []byte

func Init() {

	exe, _ := os.Executable()
	base := filepath.Dir(exe)
	path = filepath.Join(base, "URMC_DB")

	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		fmt.Println("Error creating Directory: ", err)
		return
	}

}

func createTables(db *sql.DB) error {
	fmt.Println("Creating Tables")
	data, _ := fs.ReadFile(schema, "schema/tables.sql")

	_, err := db.Exec(string(data))

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func OpenAgentDB() (*sql.DB, error) {
	return sql.Open("sqlite", fmt.Sprintf("%s/%s.db", path, global.Username))
}
