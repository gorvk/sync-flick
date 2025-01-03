package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gorvk/starterapp/internal/initializers"
)

func main() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	db := initializers.GetDBInstance()
	dbMigration(db)
}

func dbMigration(db *sql.DB) {
	fmt.Println("DB Migration Started...")

	file, err := os.ReadFile("./database/schema.sql")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(string(file))
	if err != nil {
		panic(err)
	}

	fmt.Println("DB Migration Completed!")
}
