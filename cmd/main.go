package main

import (
	"docryze-backend/app/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	print(config.Application.Name)
}
