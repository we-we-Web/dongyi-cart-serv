package infrastructure

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/we-we-Web/dongyi-cart-serv/app/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	godotenv.Load()
	dsn := os.Getenv("DB_URL")

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = DB.AutoMigrate(&entity.Cart{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	log.Println("Database connected and migrated successfully!")
	return DB, err
}
