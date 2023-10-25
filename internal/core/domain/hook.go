package domain

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// BeforeCreate func
func (t *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	logrus.Info("BeforeCreate")
	uuid, err := uuid.NewRandom() // v4
	if err != nil {
		return err
	}
	t.ID = &uuid
	return nil
}
