package db

import (
	"database/sql"
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	_ "modernc.org/sqlite"
)

func AgentDatabaseInit() error {

	fmt.Printf("%s\n", fmt.Sprintf("%s/%s.db", DBLocation, global.Username))

	db, err := sql.Open("sqlite", fmt.Sprintf("%s/%s.db", DBLocation, global.Username))

	if err != nil {
		logger.Errorf("Failed to Open/Create DB\n Error: %s", err.Error())
		return err
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		logger.Errorf("Failed to Ping DB after creation\n Error: %s", err.Error())
		return err
	}

	logger.Infof("%s - Create/Open Successfuly\n", fmt.Sprintf("%s.db", global.Username))

	createTables(db)

	return nil
}
