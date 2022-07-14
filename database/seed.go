package database

import (
	"github.com/google/uuid"
	feedbackRepo "github.com/medicplus-inc/medicplus-feedback/internal/infrastructure/repository"
	"gorm.io/gorm"
)

func seedUp() []interface{} {
	parameterID, _ := uuid.NewRandom()
	result := []interface{}{
		&feedbackRepo.FeedbackParameter{
			ID: parameterID,
		},
	}
	return result
}

func Seed(db *gorm.DB) error {
	seeders := seedUp()
	for _, seed := range seeders {
		err := db.Create(seed).Error
		if err != nil {
			return err
		}
	}
	return nil
}
