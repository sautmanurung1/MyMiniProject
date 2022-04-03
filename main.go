package main

import (
	"MyMiniProject/config"
	"MyMiniProject/routes"
)

func main(){
	config.InitDB()

	e := routes.New()
	e.Logger.Fatal(e.Start(":1234"))
}