package main

import (
	"miao_sticker_server/index"
)

var myApp *index.MyApp

func main() {
	myApp = &index.MyApp{}
	myApp.Init()
	myApp.Run()
}
