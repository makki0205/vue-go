package main

import "github.com/makki0205/template/router"

func main() {
	r := router.GetRouter()
	r.Run(":8000")
}
