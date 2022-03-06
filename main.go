package main

import "github.com/ymsht/beer-recommend-api/router"

func main() {
	e := router.Init()
	e.Logger.Fatal(e.Start(":1234"))
}
