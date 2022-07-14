package database

import (
	feedbackRepo "github.com/medicplus-inc/medicplus-feedback/internal/infrastructure/repository"
	"gorm.io/gorm"
)

func generateModels() []interface{} {
	result := []interface{}{}
	result = append(result, &feedbackRepo.FeedbackParameter{})

	return result
}

// Migrate migrates the database up
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(generateModels()...)
}
