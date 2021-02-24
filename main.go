package main

import (
	"github.com/aidar-darmenov/string.generator/model"
	"github.com/aidar-darmenov/string.generator/service"
	"github.com/gin-gonic/gin"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
	"strconv"
)

func main() {
	_cfg := model.NewConfiguration("config.json")
	_server := server(_cfg)
	_server.Run(":" + strconv.Itoa(_cfg.HTTPPort))
}

func server(_cfg *model.Configuration) *gin.Engine {
	//_db := db.NewDB(_cfg)
	//_dbm := db.NewDBManager(_db)
	//_dbm.AutoMigrate()

	_service := service.NewCoreManager( /*_db,*/ _cfg)

	_service.StartWorkers(_cfg)

	g := gin.Default()
	g.POST("/healthcheck", _service.HealthCheck)
	g.POST("/encrypt-string-list", _service.EncryptStringList)
	g.POST("/generate-string/:number", _service.GenerateString)

	return g
}
