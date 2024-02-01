package models

import (
	"time"

	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
  dsnPostgres := "postgres://postgres:postgres@db:5432"
  db, err := gorm.Open(postgres.Open(dsnPostgres), &gorm.Config{})

  err = db.AutoMigrate(&Product{})
  if err != nil {
    panic("failed to migrate")
  }

  postgresqlDB, err := db.DB()
  if err != nil {
    panic("failed to migrate")
  }

  postgresqlDB.SetMaxIdleConns(100)
  postgresqlDB.SetMaxOpenConns(100)
  postgresqlDB.SetConnMaxLifetime(time.Hour)

  DB = db
}
