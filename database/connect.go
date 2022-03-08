package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kingztech2019/blogbackend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB
func Connect()  {
	err:=godotenv.Load()
	if err != nil {
		 log.Fatal("Error load .env file")
	}
	dsn:=os.Getenv("DSN")
	database,err:=gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}else{
		log.Println("connect successfully")
	} 
	DB=database
	database.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)

	
}