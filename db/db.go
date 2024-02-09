package db

import (
	structs "Appointy/Tasks/grpc-crud/protobuf"
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDB() (*gorm.DB, error) {
	sqlDB, _ := sql.Open("pgx", "DB_URL")
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	gormDB.AutoMigrate(&structs.Task{})

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connection successful.")
	return gormDB, err

}

type DBHandler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *DBHandler {
	return &DBHandler{db}
}
