package main

import (
	"fmt"
	"github.com/habales/gotmx-play/srv"
)

func main() {
	fmt.Print("Go HTMX FTW")

	app := &srv.App{}
	app.RunApp()

}
