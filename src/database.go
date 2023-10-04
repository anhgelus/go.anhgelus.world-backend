package src

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Redirection struct {
	gorm.Model
	Slug        string
	Destination string
}

func (credentials *DatabaseCredentials) GenerateDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		credentials.Host,
		credentials.User,
		credentials.Password,
		credentials.DBName,
		credentials.Port,
	)
}

func (credentials *DatabaseCredentials) Connect() {
	var err error
	DB, err = gorm.Open(postgres.Open(credentials.GenerateDSN()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(&Redirection{})
	if err != nil {
		panic(err)
	}
}
