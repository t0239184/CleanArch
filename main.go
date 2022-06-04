package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_userDlivery "github.com/t0239184/CleanArch/app/user/delivery/http"
	_userRepo "github.com/t0239184/CleanArch/app/user/repository"
	_userUsecase "github.com/t0239184/CleanArch/app/user/usecase"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			panic(fmt.Errorf("Fatal error config file: %w \n", err))
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("Fatal error config file: %w \n", err))
		}
	}
}

func main() {
	r := gin.Default()
	db := InitDatabase()
	userRepo := _userRepo.NewUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo)
	_userDlivery.NewUserHandler(r, userUsecase)

	r.Run(":8080")
}

func InitDatabase() *gorm.DB {
	user := viper.GetString("database.connection-info.user")
	password := viper.GetString("database.connection-info.password")
	host := viper.GetString("database.connection-info.host")
	port := viper.GetInt("database.connection-info.port")
	database := viper.GetString("database.connection-info.database")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", user, password, host, port, database)
	fmt.Println(dsn)
	gormConfig := &gorm.Config{}
	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		logrus.Fatalf("[main] database.New failed: %v", err)
	}
	db.AutoMigrate()
	return db
}
