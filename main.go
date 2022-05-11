package main

import (
	"gocode/common"
	"gocode/server"
)

func init() {
	common.LoadTemplate()
}

func main() {
	server.App.Start("127.0.0.1", "8080")
}
