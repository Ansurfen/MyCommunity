package main

import "go-auth/server"

func main() {
	go server.ShutDown()
	app := server.NewApp().Default().UseMiddleWare().UseRouter()
	app.Run()
}
