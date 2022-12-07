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
	//dbname=%s
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable",
		dbSetting.Host, dbSetting.User, dbSetting.Password, dbSetting.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// err2 := db.Exec("create database " + dbSetting.Name)
	// if err2 != nil {
	// 	//handle the error
	// 	log.Fatal(err2)
	// }

	if err != nil {
		panic(err)
	}

	db.AutoMigrate()

	return &Repository{
		database: db,
	}

}

// SendSMS implements ports.IRepository
func (r *Repository) SendSMS(sms domain.SMS) error {
	if err := r.database.Create(&sms).Error; err != nil {
		return err
	}

	return nil
}

// LogSMS implements ports.IRepository
func (*Repository) LogSMS() error {
	panic("unimplemented")
}
