package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() {
	dsn := "postgres://" + os.Getenv("AZURE_PG_USER") + ":" + os.Getenv("AZURE_PG_PASSWORD") + "@hackzallopostgres.postgres.database.azure.com/postgres?sslmode=require"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("===== DB Connected =====\n")
	}
	db.Logger = db.Logger.LogMode(logger.Info)
	return
}
