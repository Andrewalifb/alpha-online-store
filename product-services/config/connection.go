package config

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Andrewalifb/alpha-online-store/product-services/entity"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	SQLDB *gorm.DB
}

type RedisConfig struct {
	RedisDB *redis.Client
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
		db.AutoMigrate(&entity.Category{}, &entity.Product{})
		return db
	}
}

func connectRedis() *redis.Client {
	redisDB := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		DB:   0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := redisDB.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Failed to connect to Redis:", err)
		return nil
	} else {
		fmt.Println("Successfully connected to Redis")
		return redisDB
	}
}

func NewConfigPostgresql() *Config {
	return &Config{
		SQLDB: connectPostgres(),
	}
}

func NewConfigRedis() *RedisConfig {
	return &RedisConfig{
		RedisDB: connectRedis(),
	}
}
