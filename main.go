package main

import "github.com/makki0205/vue-go/router"

func main() {
	r := router.GetRouter()
	r.Run(":8000")
}
