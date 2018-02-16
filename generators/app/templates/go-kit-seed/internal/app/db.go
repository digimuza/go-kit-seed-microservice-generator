package app

import (
	"fmt"

	"dev.adeoweb.biz/pas/<%= appName %>/pkg/utils"

	"github.com/go-kit/kit/log/level"
	"github.com/jinzhu/gorm"

	"github.com/go-kit/kit/log"
)

// NewDBConnection - create connection to database
func NewDBConnection(config utils.Configuration, logger log.Logger) *gorm.DB {
	level.Info(logger).Log(
		"status", "Connecting to database",
		"driver", config.GetDBDriver(),
		"host", config.GetDBHost(),
		"port", config.GetDBPort(),
		"db-name", config.GetDBName(),
	)

	url := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		config.GetDBHost(),
		config.GetDBPort(),
		config.GetDBUser(),
		config.GetDBName(),
		config.GetDBPass(),
	)

	dbconnecetion, err := gorm.Open(config.GetDBDriver(), url)
	if err != nil {
		level.Error(logger).Log("err", err)
		panic(err)
	}

	dbconnecetion.LogMode(false)

	return dbconnecetion
}
