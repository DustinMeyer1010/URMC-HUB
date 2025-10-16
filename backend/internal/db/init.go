package db

import (
	"database/sql"
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	_ "modernc.org/sqlite"
)

//go:embed schema/*.sql
var schema embed.FS
var DBConnection *sql.DB
var UserID int

func Init() {
	var err error
	dbPath := filepath.Join(getExecutableDir(), "database.db")
	fmt.Println(dbPath)
	DBConnection, err = sql.Open("sqlite", dbPath)

	if err != nil {
		panic(err)
	}

	err = createTables()
	if err != nil {
		panic(err)
	}

}

func createTables() error {

	tables, _ := schema.ReadFile("schema/init.sql")

	_, err := DBConnection.Exec(string(tables))

	return err
}

func getExecutableDir() string {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(exePath)
}

func SetUserID(agent models.Agent) {
	query := "SELECT id FROM agents WHERE username=?"

	result := DBConnection.QueryRow(query, agent.Username)

	result.Scan(&UserID)
}
