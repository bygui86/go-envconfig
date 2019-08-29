package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {

	loading()
	getting()
}

func loading() {

	viper.SetConfigType("json")
	viper.SetConfigName("config")            // name of config file (without extension)
	viper.AddConfigPath("/etc/go-envvars/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.go-envvars") // call multiple times to add many search paths
	viper.AddConfigPath(".")                 // optionally look for config in the working directory
	err := viper.ReadInConfig()              // Find and read the config file
	if err != nil {                          // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println("Config file loaded")
}

func getting() {

	env1 := viper.GetString("config1")
	fmt.Println("config1: " + env1)
}
