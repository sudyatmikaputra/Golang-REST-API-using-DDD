package config

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

			dbLogger := logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
				logger.Config{
					SlowThreshold:             time.Second, // Slow SQL threshold
					LogLevel:                  logger.Info, // Log level
					IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
					Colorful:                  true,        // Disable color
				},
			)

			dsn := getPostgresDSN()
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
				Logger: dbLogger,
			})

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
