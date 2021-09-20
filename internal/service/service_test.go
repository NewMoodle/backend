package service_test

import (
	"github.com/ZhansultanS/myLMS/backend/internal/repository"
	"github.com/ZhansultanS/myLMS/backend/internal/service"
	"github.com/ZhansultanS/myLMS/backend/pkg/database"
	"github.com/ZhansultanS/myLMS/backend/pkg/hasher"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"testing"
)

var services = service.Service{}
var pool = &pgxpool.Pool{}

func TestMain(m *testing.M) {
	if err := godotenv.Load("../../.env.local"); err != nil {
		log.Fatal("cannot load environment variables", err)
	}

	host := os.Getenv("TEST_POSTGRE_HOST")

	port, err := strconv.Atoi(os.Getenv("TEST_POSTGRE_PORT"))
	if err != nil {
		log.Fatal("cannot convert string to int", err)
	}
	username := os.Getenv("TEST_POSTGRE_USERNAME")
	password := os.Getenv("TEST_POSTGRE_PASSWORD")
	dbname := os.Getenv("TEST_POSTGRE_DATABASE")
	sslmode := os.Getenv("POSTGRE_SSLMODE")
	pools, err := strconv.Atoi(os.Getenv("POSTGRE_POOLS"))
	if err != nil {
		log.Fatal("cannot convert string to int", err)
	}

	pool, err = database.NewPostgreConnectionPool(host, port, username, password, dbname, sslmode, pools)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	hashCost, err := strconv.Atoi(os.Getenv("PASSWORD_HASH_COST"))
	if err != nil {
		log.Fatal("cannot convert string to int", err)
	}

	repositories := repository.NewRepository(pool)
	passwordhasher := hasher.BcryptHasher{Cost: hashCost}
	deps := service.Deps{
		Repositories:   repositories,
		PasswordHasher: &passwordhasher,
	}
	services = service.NewService(deps)

	os.Exit(m.Run())
}
