package main

import "GIN/routers"

var PORT = ":8080"

func main() {
	routers.StartServer().Run(PORT)
}