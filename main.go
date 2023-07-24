package main

import (
	"fmt"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_viper"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var viper *wrap_viper.WrapViper

func init() {
	viper = wrap_viper.GetViper()
	state := viper.Get("STATE")
	switch state {
	case "PRO":
		viper.LoadConfigFile("./config/", "pro")
	case "SUT":
		viper.LoadConfigFile("./config/", "sut")
	default:
		state = "DEV"
		viper.LoadConfigFile("./config/", "dev")
	}
	gin.SetMode(gin.ReleaseMode)

	fmt.Println("STATE:", state)
}

func main() {
	startHttpServer()
}

func startHttpServer() {
	r := gin.Default()
	r.Use(cors.Default())

	router(r)
	port := viper.GetInt("PORT")
	if err := r.Run(fmt.Sprintf(`:%v`, port)); err != nil {
		panic(err)
	}
}
