package db

import (
	"github.com/jinzhu/gorm"
	"github.com/ArthurHlt/microserv-helper/db/model"
	"github.com/ArthurHlt/microserv-helper/config"
	"github.com/ArthurHlt/microserv-helper/logger"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ArthurHlt/gominlog"
)
var dbInstance *gorm.DB

var loggerDb *gominlog.MinLog = logger.GetMinLog()
func GetDb() (*gorm.DB, error) {
	if dbInstance != nil {
		return dbInstance, nil
	}
	var db gorm.DB
	db, err := gorm.Open(config.GetConfig().Db.Driver, config.GetConfig().Db.DataSource)
	if err != nil {
		return nil, err
	}
	dbInstance = &db
	autoMigrate()
	return dbInstance, nil
}

func autoMigrate() {
	loggerDb.Info("Auto-migrating database")
	db, _ := GetDb()
	db.AutoMigrate(
		&model.Application{},
		&model.Instance{},
	)
}