package main

import "latihan-pertama/routers"

var PORT = ":8080"

func main() {
	routers.StartServer().Run(PORT)
}
