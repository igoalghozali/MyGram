package main

import (
	"log"

	"github.com/igoalghozali/mygram/database"
	"github.com/igoalghozali/mygram/routers"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("ENV")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Fatalln("Error loading env file:", err)
	// }

	log.Println("Env successfully loaded")
}

func main() {
	database.StartDB()
	routers.StartApp().Run(":" + viper.GetString("PORT"))
}
