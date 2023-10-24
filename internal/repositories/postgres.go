package repositories

import (
	"hugeman/internal/core/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	layoutDateTimeRFC3339 = "2006-01-02T15:04:05Z"
)

// Postgres struct
type Postgres struct {
	dbGorm *gorm.DB
}

// NewPostgres func
func NewPostgres(dbGorm *gorm.DB) *Postgres {
	logrus.Info("Migrate database ...", layoutDateTimeRFC3339)
	domain.MigrateDatabase(dbGorm)
	return &Postgres{
		dbGorm: dbGorm,
	}
}
