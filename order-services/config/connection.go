package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Andrewalifb/alpha-online-store/order-services/entity"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	SQLDB *gorm.DB
}

func connectPostgres() *gorm.DB {

	host := os.Getenv("SQL_HOST")
	port, _ := strconv.Atoi(os.Getenv("SQL_PORT"))
	user := os.Getenv("SQL_USER")
	dbname := os.Getenv("SQL_DB_NAME")
	pass := os.Getenv("SQL_PASSWORD")

	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, pass)

	db, err := gorm.Open(postgres.Open(psqlSetup), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to PostgreSQL:", err)
		return nil
	} else {
		fmt.Println("Successfully connected to PostgreSQL")
		db.AutoMigrate(&entity.PaymentMethod{}, &entity.Transaction{}, &entity.TransactionItem{})
		return db
	}
}

func NewConfigPostgresql() *Config {
	return &Config{
		SQLDB: connectPostgres(),
	}
}
