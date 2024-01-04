package main

import (
	viperConfig "github.com/asinha24/global-lib/viper-config"
	"github.com/asinha24/verify-otp/api"
	"github.com/gin-gonic/gin"
)

func init() {
	viperConfig.LoadViperConfig(".")
}

func main() {
	router := gin.Default()

	app := api.Config{Router: router}

	app.Routes()

	router.Run(":8000")

}
