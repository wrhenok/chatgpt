package main

import (
	"chatgpt3/src/controller"
	_ "chatgpt3/src/mapper"
)

func main() {
	controller.HttpServer()
}
