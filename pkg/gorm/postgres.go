package gorm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// NewPostgres
// DSN "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"

func NewPostgres(pgconf postgres.Config, gormconf *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(pgconf), gormconf)
	if err != nil {
		log.Fatalln(err)
	}
	return db, err
}
