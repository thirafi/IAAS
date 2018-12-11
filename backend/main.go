package main

import (
	"fmt"

	"./app"
	"./app/config"
)

func main() {
	fmt.Println("udah jalan coi")
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
