package controller

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Qr struct {
	ctx context.Context
}

func NewQr() *Qr {
	return &Qr{
		ctx: context.Background(),
	}
}

func (x *Qr) Startup(ctx context.Context) {
	x.ctx = ctx
}

func CreateFolder(folder_rute string) error {
	err := os.Mkdir(folder_rute, 0777)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func ImageBytes(file string) []byte {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err.Error())
	}
	return bytes
}

func MoveImageToFolder(rute string, file []byte) error {
	err := ioutil.WriteFile(rute, file, 0777)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func FileDialong(x *context.Context) string {
	file, err := runtime.OpenFileDialog(*x, runtime.OpenDialogOptions{
		DefaultDirectory:           "",
		DefaultFilename:            "",
		Title:                      "Select File",
		Filters:                    []runtime.FileFilter{{DisplayName: "Images (*.png;*.jpg)", Pattern: "*.png;*.jpg"}},
		ShowHiddenFiles:            false,
		CanCreateDirectories:       false,
		ResolvesAliases:            false,
		TreatPackagesAsDirectories: false,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	return file
}

func Path() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	return path
}

func GeneratePath(inicial string, final string) string {
	path := fmt.Sprintf("%s/%s", inicial, final)
	return path
}

func (x *Qr) ReturnImage() string {
	file := FileDialong(&x.ctx)
	if file != "" {
		// image := strings.Split(file, "/")
		// path := Path()
		// folder := GeneratePath(path, "frontend/src/assets/image/Tem")
		// image_rute := GeneratePath(folder, image[len(image)-1])
		// CreateFolder(folder)
		bytes := ImageBytes(file)
		// MoveImageToFolder(image_rute, bytes)
		return fmt.Sprintf("data:image/png;base64,%s", ToBase64(bytes)) ///GeneratePath("src/assets/image/Tem", image[len(image)-1])
	}
	return ""
}

func ImageGenerateToUrl(url string, name string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer resp.Body.Close()
	return fmt.Sprintf("data:image/png;base64,%s", ToBase64(bytes))
	// path := Path()
	// folder := GeneratePath(path, "frontend/src/assets/image/Tem")
	// CreateFolder(folder)
	// file, err := os.Create(GeneratePath(folder, name))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()

	// io.Copy(file, resp.Body)
	// return GeneratePath("src/assets/image/Tem", name)
}

func (x *Qr) ImageUrl(url string, name string) string {
	return ImageGenerateToUrl(url, name)
}

func AsyncImageUrlList(url string, name string, list *[]string, wg *sync.WaitGroup, wm *sync.RWMutex) {

	// wm.RLock()
	// resp, err := http.Get(url)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer resp.Body.Close()

	// path := Path()
	// folder := GeneratePath(path, "frontend/src/assets/image/Tem")
	// wm.RUnlock()
	// wm.Lock()
	// CreateFolder(folder)
	// file, err := os.Create(GeneratePath(folder, name))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()

	// io.Copy(file, resp.Body)

	// *list = append(*list, GeneratePath("src/assets/image/Tem", name))
	// wm.Unlock()
	// wg.Done()
	wm.RLock()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err.Error())
	}
	wm.RUnlock()
	wm.Lock()

	*list = append(*list, fmt.Sprintf("data:image/png;base64,%s", ToBase64(bytes)))
	wm.Unlock()
	wg.Done()
}

func (x *Qr) ImageUrlList(url []string, name []string) []string {
	var list []string
	wg := &sync.WaitGroup{}
	wm := &sync.RWMutex{}
	wg.Add(len(url))
	for i := 0; i < len(url); i++ {
		AsyncImageUrlList(url[i], name[i], &list, wg, wm)
	}
	wg.Wait()
	// for i := 0; i < len(url); i++ {
	// 	list = append(list, ImageGenerateToUrl(url[i], name[i]))
	// }
	return list
}
