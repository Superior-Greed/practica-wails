package main

import (
	images_organizador "changeme/backend/controller"
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx         context.Context
	organizador *images_organizador.Image
	qr          *images_organizador.Qr
}

// NewApp creates a new App application struct
func NewApp(organizador *images_organizador.Image, qr *images_organizador.Qr) *App {
	return &App{
		organizador: organizador,
		qr:          qr,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.organizador.Startup(ctx)
	a.qr.Startup(ctx)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Image() []string {
	selection, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select File",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Images (*.png;*.jpg)",
				Pattern:     "*.png;*.jpg",
			}, {
				DisplayName: "Videos (*.mov;*.mp4)",
				Pattern:     "*.mov;*.mp4",
			},
		},
	})
	if err != nil {

	}

	var list []string
	for _, route := range selection {
		list = append(list, toStringBase64(route))
	}
	// for _, element := range selection {
	// 	fmt.Println(element)
	// }
	return list
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func toStringBase64(s string) string {
	bytes, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	case "image/gif":
		base64Encoding += "data:image/gif;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)
	return base64Encoding
}
