package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var db *gorm.DB

// timer is a countdown timer that will close the database connection after 10s.
// It is reset every time the GetDatabase function is called.
var timer *time.Timer

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&Event{},
		&User{},
	); err != nil {
		return err
	}
	return nil
}

func GetDatabase() (*gorm.DB, error) {
	// reset the timer if it is not nil
	if timer != nil {
		timer.Reset(time.Second * 10)
	}

	// if the database connection is already open, return it
	if db != nil {
		return db, nil
	}

	// open the database connection
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// start a timer of 10s, after 10s close the connection
	timer = time.AfterFunc(time.Second*10, func() {
		if db == nil {
			return
		}
		sqlDB, err := db.DB()
		if err != nil {
			panic("failed to close database connection: " + err.Error())
		}
		sqlDB.Close()
		db = nil
	})

	return db, err
}
