package main

import "github.com/efydb/config"

func main() {
	config.Connect()

	CreateRouter()
}
