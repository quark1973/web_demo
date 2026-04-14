package config

import (
	"context"
	"fmt"
	"log"
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	DataBase struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		Dsn     string
	}
	Redis struct {
		Host     string
		Port     string
		Password string
		DB     int
	}
}

var AppConfig *Config

func InitConfig(ctx context.Context) (error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)	
		return err
	}
	AppConfig = &Config{}
	if err = viper.Unmarshal(AppConfig); err != nil{
		fmt.Println(err)
		return err
	}

	err=ConnectDB(ctx)
	err=InitRedis(ctx)
	if err!= nil {
		log.Fatalln(err)
	}	
	return nil
}