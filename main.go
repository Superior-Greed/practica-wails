package main

import (
	"embed"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

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

type FileLoader struct {
	http.Handler
}

func NewFileLoader() *FileLoader {
	return &FileLoader{}
}

func (h *FileLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var err error
	requestedFilename := strings.TrimPrefix(req.URL.Path, "/")

	println("Requesting file:", requestedFilename)
	fileData, err := os.ReadFile(requestedFilename)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("Could not load file %s", requestedFilename)))
	}

	res.Write(fileData)
}

// // func middlewareCarpet(next http.Handler) http.Handler {
// // 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// // 		log.Print("Executing middlewareTwo")
// // 		if r.URL.Path == "/foo" {
// // 			return
// // 		}

// // 		next.ServeHTTP(w, r)
// // 		log.Print("Executing middlewareTwo again")
// // 	})
// // }

//	func CarpetHandler(w http.ResponseWriter, r *http.Request) {
//		directorio := "./"
//		http.Handle("/", http.FileServer(http.Dir(directorio)))
//	}
//
//	func Carpet(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			fmt.Println(r.URL.Path)
//			if r.URL.Path == "/" {
//				// directorio := "./"
//				// fmt.Println("entre puta")
//				// fmt.Println(r.Host)
//				// http.Handle("/static", http.FileServer(http.Dir(directorio)))
//				//os.MkdirAll("image", 0777)
//				directorio := "./"
//				http.NewServeMux()
//				file := http.Dir(directorio)
//				http.FileServer(file)
//				http.Handle("/static", http.FileServer(http.Dir(directorio)))
//				//http.ListenAndServe(":8080", nil)
//				next.ServeHTTP(w, r)
//			}
//			next.ServeHTTP(w, r)
//		})
//	}
func main() {
	// Create an instance of the app structure

	organizador := controller.NewImage()
	qr := controller.NewQr()
	app := NewApp(organizador, qr)

	path, err_path := os.Getwd()
	if err_path != nil {
		fmt.Println(err_path)
	}

	os.RemoveAll(fmt.Sprintf("%s/image", path))
	os.MkdirAll(fmt.Sprintf("%s/image/tem", path), 0777)
	os.MkdirAll(fmt.Sprintf("%s/build", path), 0777)

	icon, problem := ioutil.ReadFile(path + "/build/appicon.png")
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
			Assets:  assets,
			Handler: NewFileLoader(),
			//Middleware: assetserver.ChainMiddleware(Carpet),
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
