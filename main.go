package main

import "beer-recommend-api/router"

func main() {
	e := router.Init()
	e.Logger.Fatal(e.Start(":1234"))
}
