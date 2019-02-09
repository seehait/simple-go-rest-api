package main

import (
	"simple-go-rest-api/src/servers"
)

func main() {
	// Initialization
	e := servers.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
