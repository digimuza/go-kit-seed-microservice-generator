package main

import (
	"os"

	internalapp "dev.adeoweb.biz/pas/<%= appName %>/internal/app"
	utils "dev.adeoweb.biz/pas/<%= appName %>/pkg/utils"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	var config utils.Configuration
	{
		config = utils.NewConfiguration(
			os.Getenv("TITLE"),
			os.Getenv("ORG_NAME"),
			os.Getenv("APP_NAME"),
			os.Getenv("DB_DRIVER"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			false,
			os.Getenv("SSL_CERT_PATH"),
			os.Getenv("SSL_KEY_PATH"),
			os.Getenv("DOMAIN"),
			os.Getenv("PORT"),
			os.Getenv("LOG_LEVEL"),
			os.Getenv("LOG_ENABLED"),
			os.Getenv("ZIPKIN_HOST"),
			os.Getenv("ZIPKIN_PORT"),
			os.Getenv("PROMETHEUS_METRICS_PORT"),
		)
	}
	internalapp.Start(config)
}
