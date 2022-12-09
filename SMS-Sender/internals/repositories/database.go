package repositories

import (
	"fmt"

	"github.com/alishokri1661s/SMS-Gateway/SMS-Sender/internals/core/ports"
	"github.com/alishokri1661s/SMS-Gateway/Utils/conf"
	"github.com/alishokri1661s/SMS-Gateway/Utils/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	database *gorm.DB
}

var _ ports.IRepository = (*Repository)(nil)

func CreateDabaseConnection() *Repository {
	dbSetting := conf.DatabaseSetting
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbSetting.Host, dbSetting.User, dbSetting.Password, dbSetting.Name, dbSetting.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &Repository{
		database: db,
	}

}

func (r *Repository) UpdateSmsStatus(sms models.SMS, status models.MessageStatus) (models.SMS, error) {
	sms.Status = status
	err := r.database.Save(&sms).Error
	return sms, err
}
