package main

import (
	"demoProject4mall/conf"
	"demoProject4mall/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
