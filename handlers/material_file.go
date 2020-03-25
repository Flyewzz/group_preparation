package handlers

import (
	// "log"

	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// // MaterialFiles
// r.HandleFunc("/material/{id}/files", hd.AddMaterialFilesHandler).Methods("POST")
// r.HandleFunc("/material/{id}/files", hd.GetMaterialFilesHandler).Methods("GET")
// r.HandleFunc("/material/file/downloading", hd.MaterialFileDownloadHandler).Methods("GET")
// r.HandleFunc("/material/{id}/files/downloading", hd.MaterialFilesDownloadHandler).Methods("GET")

func (hd *HandlerData) GetMaterialFilesHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	strId := mux.Vars(r)["id"]
	materialId, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	files, err := hd.MaterialController.GetMaterialFileController().GetAll(materialId)

	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(files)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (hd *HandlerData) MaterialFileDownloadHandler(w http.ResponseWriter, r *http.Request) {
	strId := r.URL.Query().Get("id")
	if strId == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if id < 1 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	file, err := hd.MaterialController.GetMaterialFileController().GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	openedFile, err := os.Open(file.Path)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer openedFile.Close()
	//File is found, create and send the correct headers

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	openedFile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := openedFile.Stat()                   //Get info from the file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string
	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+file.Name)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	openedFile.Seek(0, 0)
	io.Copy(w, openedFile) //'Copy' the file to the client
	data, err := ioutil.ReadAll(openedFile)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

// *Not realized yet

// func (hd *HandlerData) MaterialFilesDownloadHandler(w http.ResponseWriter, r *http.Request) {

// }

// func (hd *HandlerData) AddMaterialFileHandler(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		http.Error(w, "Bad Request", http.StatusBadRequest)
// 		return
// 	}
// 	strId := mux.Vars(r)["id"]
// 	subjectId, err := strconv.Atoi(strId)
// 	if err != nil {
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 		return
// 	}
// 	params := r.PostFormValue
// 	name := params("name")
// 	strTypeId := params("type_id")
// 	typeId, err := strconv.Atoi(strTypeId)
// 	if err != nil {
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 		return
// 	}
// 	// AuthorId always will be 1 in debug (the first user makes everything temporary)
// 	addedId, err := hd.MaterialController.Add(subjectId, name, typeId, 1)
// 	if err != nil {
// 		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
// 		return
// 	}
// 	w.Write([]byte(fmt.Sprintf("Added with id %d\n", addedId)))
// }

// func (hd *HandlerData) MaterialByIdRemoveHandler(w http.ResponseWriter, r *http.Request) {
// 	strId := r.URL.Query().Get("id")
// 	id, err := strconv.Atoi(strId)
// 	if err != nil {
// 		http.Error(w, "Bad Request", http.StatusBadRequest)
// 		return
// 	}
// 	if id < 1 {
// 		http.Error(w, "Bad Request", http.StatusBadRequest)
// 		return
// 	}
// 	err = hd.MaterialController.RemoveById(id)
// 	if err != nil {
// 		if err == errs.MaterialDoesntExist {
// 			http.Error(w, "Not found", http.StatusNotFound)
// 		} else {
// 			http.Error(w, "Server Internal Error", http.StatusInternalServerError)
// 		}
// 		return
// 	}
// 	w.Write([]byte(fmt.Sprintf("A material with id %d was deleted\n", id)))
// }

// func (hd *HandlerData) MaterialByIdGetHandler(w http.ResponseWriter, r *http.Request) {
// 	strId := r.URL.Query().Get("id")
// 	id, err := strconv.Atoi(strId)
// 	if err != nil {
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 		return
// 	}
// 	if id < 1 {
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 		return
// 	}
// 	material, err := hd.MaterialController.GetById(id)
// 	if err != nil {
// 		switch err {
// 		case sql.ErrNoRows:
// 			http.Error(w, "Not found", http.StatusNotFound)
// 		default:
// 			http.Error(w, "Server Internal Error", http.StatusInternalServerError)
// 		}
// 		return
// 	}
// 	data, _ := json.Marshal(material)
// 	w.Write(data)
// }

// func (hd *HandlerData) AllMaterialsRemoveHandler(w http.ResponseWriter, r *http.Request) {
// 	strId := mux.Vars(r)["id"]
// 	id, err := strconv.Atoi(strId)
// 	if err != nil {
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 		return
// 	}
// 	err = hd.MaterialController.RemoveAll(id)
// 	if err != nil {
// 		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
// 		return
// 	}
// 	w.Write([]byte(fmt.Sprintf("All materials was deleted for the subject with id %d\n", id)))
// }
