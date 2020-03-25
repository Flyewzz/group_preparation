package handlers

import (
	// "log"

	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
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

func (hd *HandlerData) MaterialFilesDownloadHandler(w http.ResponseWriter, r *http.Request) {
	reader, err := r.MultipartReader()
	if err != nil {
		fmt.Println(err)
	}
	var uuids []string
	for {
		file, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		if file.FileName() == "" {
			continue
		}
		// Generate UUID key as a filename to store it into the temporary folder
		uuid := uuid.NewV4().String()
		dst, err := os.Create(viper.GetString("") + uuid)
		if err != nil {
			fmt.Println(err)
		}
		uuids = append(uuids, uuid)

		io.Copy(dst, file)
	}

}

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
