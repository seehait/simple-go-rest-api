package main

import (
	"simple-go-rest-api/src/databases"
	"simple-go-rest-api/src/servers"
)

func main() {
	// Initialization
	databases.Init()
	e := servers.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
