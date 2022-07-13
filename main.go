package main

import (
	"MyCommunity/server"
)

func main() {
	go server.ShutDown()
	app := server.NewApp().Default().UseRouter().UseMiddleWare()
	app.Run()
}
