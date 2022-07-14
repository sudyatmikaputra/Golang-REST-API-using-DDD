package config

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var once sync.Once

func MockDB(dbMock *gorm.DB) {
	db = dbMock
}

func DB() *gorm.DB {
	once.Do(func() {
		if db == nil {
			var err error

			dsn := getPostgresDSN()
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

			if nil != err {
				GetLogger().WithFields(logrus.Fields{
					"dsn":       dsn,
					"dialector": "postgres",
				}).Fatal("Failed to create DB Connection ", err.Error())
			}
		}
	})

	return db
}

func getPostgresDSN() string {
	return fmt.Sprintf(GetValue(DATABASE_CONNECTION_STRING), GetValue(DATABASE_HOST), GetValue(DATABASE_USER), GetValue(DATABASE_PASS), GetValue(DATABASE_NAME), GetValue(DATABASE_PORT), GetValue(DATABASE_SSL))
}
