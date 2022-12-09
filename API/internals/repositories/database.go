package repositories

import (
	"fmt"

	"github.com/alishokri1661s/SMS-Gateway/API/internals/core/ports"
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

	db.AutoMigrate(&models.SMS{})

	return &Repository{
		database: db,
	}

}

// SendSMS implements ports.IRepository
func (r *Repository) SendSMS(sms models.SMS) (models.SMS, error) {
	err := r.database.Create(&sms).Error
	return sms, err
}

// LogSMS implements ports.IRepository
func (r *Repository) LogSMS() ([]models.SMS, error) {
	allsms := []models.SMS{}
	err := r.database.Find(&allsms).Error
	return allsms, err
}
