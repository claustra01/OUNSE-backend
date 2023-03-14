package database

import "gorm.io/gorm"

func Close(db *gorm.DB) {
	if sqlDB, err := db.DB(); err != nil {
		panic(err)
	} else {
		if err := sqlDB.Close(); err != nil {
			panic(err)
		}
	}
}
