package controller

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"changeme/backend/db"
	response "changeme/backend/models"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm"
)

type ImagesRouteBase64 struct {
	Base64        string `json:"base64,omitempty"`
	Url           string `json:"url"`
	UrlServer     string `json:"url_server"`
	UrlServerPath string `json:"url_server_path"`
}

type Image struct {
	ctx        context.Context
	Image      *ImagesRouteBase64   `json:"data"`
	Image_list *[]ImagesRouteBase64 `json:"data_list,omitempty"`
}

func NewImage() *Image {
	return &Image{
		ctx: context.Background(),
	}
}

func (x *Image) Startup(ctx context.Context) {
	x.ctx = ctx
}

func ToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func ToStringBase64(s string, wg *sync.WaitGroup, m *sync.RWMutex, list *[]ImagesRouteBase64) {

	m.RLock()
	bytes, err := ioutil.ReadFile(s)
	dates := strings.Split(s, "/")

	if err != nil {
		log.Fatal(err)
	}
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	os.Mkdir(path+"/image", 0777)
	ioutil.WriteFile(fmt.Sprintf(path+"/image/tem-%s", dates[len(dates)-1]), bytes, 0777)
	if err != nil {
		log.Println(err)
	}

	m.RUnlock()
	m.Lock()
	var base64Encoding string

	mimeType := http.DetectContentType(bytes)

	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	case "image/gif":
		base64Encoding += "data:image/gif;base64,"
	}

	base64Encoding += ToBase64(bytes)

	*list = append(*list, ImagesRouteBase64{
		Base64:    base64Encoding,
		Url:       s,
		UrlServer: fmt.Sprintf("/image/tem-%s", dates[len(dates)-1]),
		//UrlServerPath: fmt.Sprintf(path+"/frontend/src/assets/image/tem-%s", dates[len(dates)-1]),
	})
	m.Unlock()

	defer wg.Done()
}

func (x *Image) Images() []ImagesRouteBase64 {

	wg := &sync.WaitGroup{}
	mg := &sync.RWMutex{}
	selection, err := runtime.OpenMultipleFilesDialog(x.ctx, runtime.OpenDialogOptions{
		DefaultDirectory:           "",
		DefaultFilename:            "",
		Title:                      "Select File",
		Filters:                    []runtime.FileFilter{{DisplayName: "Images (*.png;*.jpg;*.gif)", Pattern: "*.png;*.jpg;*.gif"}, {DisplayName: "Videos (*.mov;*.mp4)", Pattern: "*.mov;*.mp4"}},
		ShowHiddenFiles:            false,
		CanCreateDirectories:       false,
		ResolvesAliases:            false,
		TreatPackagesAsDirectories: false,
	})
	if err != nil {
		log.Fatal(err)
	}

	if len(selection) > 0 {
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}

		os.RemoveAll(path + "/image")
	}
	list := &[]ImagesRouteBase64{}

	wg.Add(len(selection))
	for _, route := range selection {
		go ToStringBase64(route, wg, mg, list)
	}
	wg.Wait()
	fmt.Println(len(*list))
	return *list
}

func (x *Image) Carpet() string {

	selection, err := runtime.OpenDirectoryDialog(x.ctx, runtime.OpenDialogOptions{
		DefaultDirectory:           "",
		DefaultFilename:            "",
		Title:                      "Select File",
		ShowHiddenFiles:            false,
		CanCreateDirectories:       false,
		ResolvesAliases:            false,
		TreatPackagesAsDirectories: false,
	})
	if err != nil {
		log.Fatal(err)
	}

	return selection
}

func (x *Image) TransferImageToFolder(carpet string, initial_rute string, final_rute string) response.JsonResponse {
	if strings.Trim(carpet, " ") != "" && strings.Trim(initial_rute, " ") != "" && strings.Trim(final_rute, " ") != "" {
		carpet_rute := fmt.Sprintf("%s/%s", final_rute, carpet)
		os.Mkdir(carpet_rute, 0777)
		files, err := ioutil.ReadFile(initial_rute)
		dates := strings.Split(initial_rute, "/")
		if err != nil {
			log.Fatal(err)
			return response.JsonResponse{Text: "Problema con la ruta inicial", Value: false}
		}
		err_insert := ioutil.WriteFile(fmt.Sprintf("%s/%s", carpet_rute, dates[len(dates)-1]), files, 0777)
		if err_insert != nil {
			log.Fatal(err_insert)
			return response.JsonResponse{Text: "Problema al transferir la imagen a la nueva carpeta", Value: false}
		}
		err_remove := os.Remove(initial_rute)
		if err_remove != nil {
			log.Fatal(err_remove)
		}
		return response.JsonResponse{Text: "Transferido correctamente", Value: true}
	}
	return response.JsonResponse{Text: "Un valor es nulo", Value: false}
}

// --- tarea pendiente
func DirFiles(folderdir string, files_list *[]string) {
	files, err := ioutil.ReadDir(folderdir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			DirFiles(fmt.Sprintf("%s/%s", folderdir, file.Name()), files_list)
		} else {
			*files_list = append(*files_list, fmt.Sprintf("%s/%s", folderdir, file.Name()))
		}
	}
}

func RedirecImages(db *gorm.DB, file string, dir_rute string, wg *sync.WaitGroup, wm *sync.RWMutex) {
	wm.Lock()
	var split_path = strings.Split(file, "/")
	var text = strings.Split(split_path[len(split_path)-1], "_")
	folder := response.Folder{}
	db.Select("rute").First(&folder, "termination_image=?", text[0]+"_")
	fmt.Println(folder.Rute)
	if strings.Trim(folder.Rute, " ") != "" {
		image, err := ioutil.ReadFile(file)
		os.Mkdir(fmt.Sprintf("%s/%s", dir_rute, folder.Rute), 0777)
		if err != nil {
			log.Fatalln(err)
		}
		ioutil.WriteFile(fmt.Sprintf("%s/%s/%s", dir_rute, folder.Rute, split_path[len(split_path)-1]), image, 0777)
	}
	wm.Unlock()

	wg.Done()
}

func (x *Image) TransferImages(folder_init string, folder_final string) {
	var listdir []string
	var files_list []string
	wg := &sync.WaitGroup{}
	wm := &sync.RWMutex{}
	DB := db.DataBase{}.DbConnection()

	files, err := ioutil.ReadDir(folder_init)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			listdir = append(listdir, fmt.Sprintf("%s/%s", folder_init, file.Name()))
		} else {
			files_list = append(files_list, fmt.Sprintf("%s/%s", folder_init, file.Name()))
		}
	}
	for _, dir := range listdir {
		DirFiles(dir, &files_list)
	}
	wg.Add(len(files_list))
	for _, file := range files_list {
		RedirecImages(DB, file, folder_final, wg, wm)
	}
	wg.Wait()
}

// ----
func (x *Image) AddFolderImage(folder response.Folder) response.Folder {
	DB := db.DataBase{}.DbConnection()
	Folder := response.Folder{Name: folder.Name, Description: folder.Description, Rute: folder.Rute, TerminationImage: folder.TerminationImage}
	DB.Create(&Folder)
	fmt.Println(Folder)
	return Folder
}

func (x *Image) AllFolderImage() []response.Folder {
	DB := db.DataBase{}.DbConnection()
	Folder := []response.Folder{}
	DB.Find(&Folder)
	return Folder
}
