package db

import (
	"database/sql"
	"fmt"

	pb "Hero/Tasks/graphql-crud/protobuf"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDB() (*gorm.DB, error) {
	sqlDB, _ := sql.Open("pgx", "DB_URL")
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic("Unable to connect to DB")
	}
	gormDB.AutoMigrate(&pb.Task{})
	fmt.Println("DB connection successful")
	return gormDB, err
}

type DBHandler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *DBHandler {
	return &DBHandler{db}
}
