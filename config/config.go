package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// LoadENV :nodoc:
func LoadENV() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(fmt.Sprintf("ERROR: %s", err))
	}

	return err
}

// ENV :nodoc:
func ENV() string {
	return os.Getenv("ENV")
}

// HTTPPort :nodoc:
func HTTPPort() string {
	return os.Getenv("PORT")
}

// DBHost :nodoc:
func DBHost() string {
	return os.Getenv("DB_HOST")
}

// DBUser :nodoc:
func DBUser() string {
	return os.Getenv("DB_USERNAME")
}

// DBPassword :nodoc:
func DBPassword() string {
	return os.Getenv("DB_PASSWORD")
}

// DBName :nodoc:
func DBName() string {
	return os.Getenv("DB_DATABASE")
}

// DBPort :nodoc:
func DBPort() string {
	return os.Getenv("DB_PORT")
}

// RedisHost :nodoc:
func RedisHost() string {
	return os.Getenv("REDIS_HOST")
}

// RedisPort :nodoc:
func RedisPort() string {
	return os.Getenv("REDIS_PORT")
}

// RedisDB :nodoc:
func RedisDB() int {
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	return db
}

// RedisMaxIdle :nodoc:
func RedisMaxIdle() int {
	mi, _ := strconv.Atoi(os.Getenv("REDIS_MAX_IDLE"))
	return mi
}

// RedisMaxActive :nodoc:
func RedisMaxActive() int {
	ma, _ := strconv.Atoi(os.Getenv("REDIS_MAX_ACTIVE"))
	return ma
}

// RedisTTL :nodoc:
func RedisTTL() int {
	ttl, _ := strconv.Atoi(os.Getenv("REDIS_TTL"))
	return ttl
}
