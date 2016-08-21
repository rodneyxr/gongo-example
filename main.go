package main

import (
	"fmt"

	"github.com/rodneyxr/gongo"
)

func main() {
	fmt.Println("Listening on port 8000")
	app := gongo.New()
	app.RoutesFunc(Routes)
	app.Run()
}
