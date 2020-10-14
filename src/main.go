package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"go-web/src/config"
	"go-web/src/model"
	"go-web/src/router"

	"errors"
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

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Printf("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Printf(http.ListenAndServe(viper.GetString("addr"), engine).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Println("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
