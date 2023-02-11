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
	viper.LoadConfigFile("../config", "dev")
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
