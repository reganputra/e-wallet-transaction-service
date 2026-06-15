package helpers

import (
	"e-wallet-transaction-service/internal/models"
	"fmt"
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func SetupMySql() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", GetEnv("DB_USER", ""), GetEnv("DB_PASSWORD", ""), GetEnv("DB_HOST", "127.0.0.1"), GetEnv("DB_PORT", "3306"), GetEnv("DB_NAME", ""))

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}
	logrus.Info("Database Successfully Connected!")

	DB.AutoMigrate(&models.Transaction{})

	DB.Logger = logger.Default.LogMode(logger.Info)
}
