package main

import (
	// "image/color"
	"log"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/layout"
	"github.com/Ayr2017/rigger/internal/initvar"
	"github.com/Ayr2017/rigger/internal/checkdocker"
	"github.com/Ayr2017/rigger/internal/dockercommands/showimages"
)

func main() {
	fmt.Println("Run application!")
	conf, err := initvar.Init()
	var label3Text string

	if err!= nil {
		log.Fatal(err)
	}

	a := app.New()
	w := a.NewWindow(conf.AppName)

	dockerCheck := checkdocker.CheckDocker()

	if dockerCheck == true {
		label3Text = "All Right"
    } else {
		label3Text = "You have not Docker"
	}

	result := showimages.GetAll()
	fmt.Println(result[1].Repository)

	w.Resize(fyne.NewSize(640, 480))
	label1 := widget.NewLabel("Menu!")
	label2 := widget.NewLabel(label3Text)
	label3 := widget.NewLabel("Hello World!")

	grid := container.NewGridWithColumns(3, label1, label2, label3)

	w.SetContent(grid)
	w.ShowAndRun()
}
