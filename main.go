package main

import (
	"fmt"
	"github.com/nurchulis/idempotency/internal/bootstrap"
	"github.com/nurchulis/idempotency/internal/handler"
	"github.com/nurchulis/idempotency/internal/repositories"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/nurchulis/idempotency/config"
	"github.com/nurchulis/idempotency/internal/ucase/v1/user"
)

func main() {
	if errEnv := godotenv.Load(); errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	conf := config.Initialize()

	redis := bootstrap.RegistryRedis(conf.RedisConfig)

	db, err := bootstrap.RegistrySQL(conf.DBConfig)
	if err != nil {
		log.Fatal(err)
	}

	if conf.DBConfig.AutoMigrate {
		if err := bootstrap.AutoMigrate(db); err != nil {
			log.Fatal(err)
		}
	}

	if conf.DBConfig.AutoSeed {
		if err := bootstrap.AutoSeed(db); err != nil {
			log.Fatal(err)
		}
	}

	userRepo := repositories.NewUserRepo(db)

	userList := user.NewUserList(redis, userRepo)

	router := mux.NewRouter()
	// path: localhost:8000/api/v1/users
	// header: X-Request-ID
	apiV1 := router.PathPrefix("/api/v1").Subrouter()
	apiV1.HandleFunc("/users", handler.Http(userList)).Methods(http.MethodGet)

	log.Println(fmt.Sprintf("Server Running on port %d", conf.AppPort))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.AppPort), router))
}
