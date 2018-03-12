package main

import "github.com/chotchy-inc/PATRAProductAPI/router"

func main() {
	r := router.GetRouter()
	r.Run(":8000")
}
