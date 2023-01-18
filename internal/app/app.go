package app

import (
	"context"
	"fmt"
	"log"
	"os"

	"echo-mysql/internal/handler"
	"echo-mysql/internal/service"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ApplicationContext struct {
	UserHandler *handler.UserHandler
}

func NewApp(ctx context.Context) (*ApplicationContext, error) {
	var DB *gorm.DB
	var _ = godotenv.Load(".env")
	var ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("user"),
		os.Getenv("pass"),
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("db"))

	var err error
	DB, err = gorm.Open(mysql.Open(ConnectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	userService := service.NewUserService(DB)
	userHandler := handler.NewUserHandler(userService)

	return &ApplicationContext{
		UserHandler: userHandler,
	}, nil
}
