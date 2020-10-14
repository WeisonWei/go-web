package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go-web/src/model"

	"go-web/src/config"
	"go-web/src/router"

	"log"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiServer config file path.")
)

func main() {

	// init config
	pflag.Parse()

	// 具体的配置可传参 ./apiserver -c config.yaml
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// init db
	model.DB.Init()
	defer model.DB.Close()

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	engine := gin.New()

	// gin middlewares
	var middlewares []gin.HandlerFunc

	// Routes.
	router.Load(
		// Cores.
		engine,

		// Middlewares.
		middlewares...,
	)
	port := viper.GetString("addr")
	log.Printf("Start to listening the incoming requests on http address: %s", port)
	//log.Printf(http.ListenAndServe(":8080", engine).Error())
	log.Printf(http.ListenAndServe(port, engine).Error())
	time.Sleep(time.Second * 3)
	log.Printf("it is time to ping")
}
