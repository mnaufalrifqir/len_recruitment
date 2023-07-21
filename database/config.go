package database

import (
	"len_recruitment/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {

	dsn := "host=localhost user=postgres password=naufalrifqi13 dbname=len_recruitment port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&model.Buku{}, &model.Mahasiswa{}, &model.DataPeminjaman{})
}
