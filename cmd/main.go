package main

import (
	"log"
	"os"

	"github.com/EltIsma/YandexLavka/cmd/internal/controllers"
	"github.com/EltIsma/YandexLavka/cmd/internal/repository"
	httpSer "github.com/EltIsma/YandexLavka/cmd/internal/servers/http"
	"github.com/EltIsma/YandexLavka/cmd/internal/services"
	"github.com/EltIsma/YandexLavka/pkg/db/postgresql"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil{
		logrus.Fatalf("error: %s", err.Error())
	}
    
	if err:= godotenv.Load(); err != nil{
		logrus.Fatalf("error loading env variables %s", err.Error())
	}

	db, err := postgresql.NewPostgresDBClient(postgresql.Config{
		Host: viper.GetString("db.host"),
		//Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname: viper.GetString("db.dbname"),
		//SSLMode: viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := services.NewService(repos)
 	handlers := controllers.NewHandler(services)
	srv := new(httpSer.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running server: %s", err.Error())
	}


	
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
 }
