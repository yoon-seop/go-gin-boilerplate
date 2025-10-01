package config

import (
	"fmt"
	"go-gin-boilerplate/internal/app/post"
	"go-gin-boilerplate/internal/app/user"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func DatabaseInitialize(conf *AppConfig) error {
	var logLevel logger.LogLevel
	switch conf.Env() {
	case Development, Staging:
		logLevel = logger.Info
	default:
		logLevel = logger.Error
	}

	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logLevel,
			Colorful: true,
		},
	)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=Local",
		conf.Database.Username,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name,
		conf.Database.Charset,
		conf.Database.ParseTime,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return err
	}

	// if err = db.AutoMigrate(
	// 	&entity.User{},
	// 	&entity.Post{},
	// ); err != nil {
	// 	log.Fatalf("DB AutoMigrate failed: %v", err)
	// }

	user.InitRepository(db)
	post.InitRepository(db)

	DB = db

	return nil
}
