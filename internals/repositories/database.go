package repositories

import (
	"fmt"

	"github.com/alishokri1661s/SMS-Gateway/conf"
	"github.com/alishokri1661s/SMS-Gateway/internals/core/domain"
	"github.com/alishokri1661s/SMS-Gateway/internals/core/ports"
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

	db.AutoMigrate(&domain.SMS{})

	return &Repository{
		database: db,
	}

}

// SendSMS implements ports.IRepository
func (r *Repository) SendSMS(sms domain.SMS) (domain.SMS, error) {
	err := r.database.Create(&sms).Error
	return sms, err
}

// LogSMS implements ports.IRepository
func (r *Repository) LogSMS() ([]domain.SMS, error) {
	allsms := []domain.SMS{}
	err := r.database.Find(&allsms).Error
	return allsms, err
}
