package main

import (
	"log"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/Ayr2017/rigger/internal/initvar"
)

func main() {
	fmt.Println("Hello World!")
	conf, err := initvar.Init()

	if err!= nil {
		log.Fatal(err)
	}

	a := app.New()
	w := a.NewWindow(conf.AppName)

	w.Resize(fyne.NewSize(640, 480))


	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
