package main

import "im/router"

func main() {
	e := router.Router()
	e.Run("localhost:8080")
}
