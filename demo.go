package main

import (
	"gin-demo/conf"
	"gin-demo/server"
)

func main() {
	conf.Init()

	r := server.NewRouter()
	err := r.Run(":3001")
	if err != nil {
		return
	}
}
