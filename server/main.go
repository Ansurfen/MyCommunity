package main

import (
	"MyCommunity/server"
)

func main() {
	go server.ShutDown()
	app := server.NewApp().Default().UseMiddleWare().UseRouter()
	app.Run()
}
