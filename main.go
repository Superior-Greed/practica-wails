package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	"io/ioutil"

	controller "changeme/backend/controller"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

//go:embed all:frontend/dist
var assets embed.FS
var icon []byte

func main() {
	// Create an instance of the app structure
	path, err_path := os.Getwd()
	if err_path != nil {
		fmt.Println(err_path)
	}
	os.RemoveAll(path + "/frontend/src/assets/image")
	os.Mkdir(path+"/frontend/src/assets/image", 0777)
	organizador := controller.NewImage()
	qr := controller.NewQr()
	app := NewApp(organizador, qr)
	icon, problem := ioutil.ReadFile("./build/Greed3.png")
	if problem != nil {
		log.Fatal(problem)
	}

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "organizador",
		Width:     1024,
		Height:    768,
		MinWidth:  800,
		MinHeight: 500,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour:   &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:          app.startup,
		LogLevel:           logger.DEBUG,
		LogLevelProduction: logger.ERROR,
		Bind: []interface{}{
			app,
			organizador,
			qr,
		},
		Linux: &linux.Options{
			Icon:                icon,
			WindowIsTranslucent: false,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func Assets() {
	panic("unimplemented")
}
