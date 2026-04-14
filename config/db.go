package config

import (
	"context"
	"demo/global"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
func ConnectDB(ctx context.Context) error {
	dsn := AppConfig.DataBase.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	sqldb,err:=db.DB()
	
	//设置连接池
	sqldb.SetMaxIdleConns(10)
	sqldb.SetMaxOpenConns(100)
	sqldb.SetConnMaxLifetime(time.Hour)

	if err!= nil {
		return err
	}
	global.Db=db
	
	log.Println("mysql connected")
	return nil
}