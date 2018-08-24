package main

import (
	"github.com/makki0205/config"
	"github.com/makki0205/vue-go/router"
)

func main() {
	r := router.GetRouter()
	r.Run(config.Env("port"))
}
