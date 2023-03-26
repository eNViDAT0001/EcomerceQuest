package main

import (
	"fmt"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_viper"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

var viper *wrap_viper.WrapViper

func init() {
	os.Setenv("STATE", "state.(string)")

	viper = wrap_viper.GetViper()
	state := viper.Get("STATE")
	if state == "PRODUCTION" {
		viper.LoadConfigFile("./config/", "prod")
	} else if state == "STAGING" {
		viper.LoadConfigFile("./config/", "staging")
	} else {
		viper.LoadConfigFile("./config/", "dev")
	}
	gin.SetMode(gin.ReleaseMode)

	fmt.Println("STATE:", state)
}

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	router(r)

	port := viper.GetInt("PORT")
	if err := r.Run(fmt.Sprintf(`:%v`, port)); err != nil {
		panic(err)
	}
}
