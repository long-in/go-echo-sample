package main

import (
	"github.com/long-in/go-echo-sample/app"
)

// main entry
func main() {
	// init server
	app.Init()

	// run server
	app.Server.Logger.Fatal(app.Server.Start(":80"))
}
