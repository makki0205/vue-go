package main

import "github.com/aykikr/recipo/router"

func main() {
	r := router.GetRouter()
	r.Run(":8000")
}
