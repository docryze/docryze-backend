package main

import (
	"docryze-backend/internal/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	print(config.Application.Name)
}
