package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/test-tecnico-grupo-lagos/backend/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() error {
	godotenv.Load()
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBNAME"))
	var err error
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println("!!!ERROR CONECTANDO LA BBDD -> ", err.Error())
	}
	db.AutoMigrate(&entities.User{}, &entities.Cart{})
	fmt.Println("CONEXION EXITOSA A BBDD")
	return nil
}

func GetInstance() *gorm.DB {
	return db
}
