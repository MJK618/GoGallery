package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"strings"

	"fyne.io/fyne/v2/app"
	//"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"log"
	//"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	rootdir := "images"
	files, err := ioutil.ReadDir(rootdir)
	if err != nil {
		log.Fatal(err)
	}
	var imageArray []string
	tabs := container.NewAppTabs()
	for _, file := range files {
		//fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir() {
			fileExtension := strings.Split(file.Name(), ".")[1]
			if fileExtension == "png" || fileExtension == "jpeg" || fileExtension == "jpg" {
				imageArray = append(imageArray)
				image := canvas.NewImageFromFile(rootdir + "\\" + file.Name())
				//image.FillMode = canvas.ImageFillOriginal
				tabs.Append(container.NewTabItem(file.Name(), image))
			}
		}
	}
	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(tabs)
	w.Resize(fyne.NewSize(500, 500))
	w.ShowAndRun()
}
