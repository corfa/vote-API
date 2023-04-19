package main

import (
	"Voting-API/app"
)

func main() {

	application := &app.App{}
	application.Initialize()
	application.Run(":3000")

}
