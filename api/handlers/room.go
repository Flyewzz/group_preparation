package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Flyewzz/group_preparation/auth"
	"github.com/gorilla/mux"
)

func (hd *HandlerData) GetRoomHandler(w http.ResponseWriter, r *http.Request) {
	strId := r.URL.Query().Get("id")
	roomId, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if roomId < 1 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	userData := r.Context().Value("user_claims").(auth.Claims)
	banned, err := hd.RoomController.IsBanned(userData.UserId, roomId)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if banned {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	room, err := hd.RoomController.GetById(roomId)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(room)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (hd *HandlerData) AddRoomHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	strId := mux.Vars(r)["id"]
	subjectId, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if subjectId < 1 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	params := r.PostFormValue
	strTypeId := params("type_id")
	typeId, err := strconv.Atoi(strTypeId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if typeId < 1 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	name := params("name")
	userData := r.Context().Value("user_claims").(auth.Claims)
	roomId, uuid, err := hd.RoomController.Add(
		name,
		subjectId,
		typeId,
		userData.UserId,
	)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	type temp struct {
		Id   int    `json:"room_id"`
		UUID string `json:"uuid"`
	}
	room := temp{
		Id:   roomId,
		UUID: uuid,
	}
	data, err := json.Marshal(room)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (hd *HandlerData) GetRoomsHandler(w http.ResponseWriter, r *http.Request) {
	userData := r.Context().Value("user_claims").(auth.Claims)
	rooms, err := hd.RoomController.GetAll(userData.UserId)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(rooms)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (hd *HandlerData) JoinRoomHandler(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]
	if uuid == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	userData := r.Context().Value("user_claims").(auth.Claims)
	err := hd.RoomController.Join(userData.UserId, uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func (hd *HandlerData) BanRoomHandler(w http.ResponseWriter, r *http.Request) {
	strId := mux.Vars(r)["id"]
	roomId, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if roomId < 1 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	authorData := r.Context().Value("user_claims").(auth.Claims)
	expectedAuthorId, err := hd.RoomController.GetAuthorId(roomId)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Only the author of the room can ban/unban other users
	// Check permissions
	if authorData.UserId != expectedAuthorId {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	params := r.PostFormValue

	strBlockId := params("user_id")
	blockingUserId, err := strconv.Atoi(strBlockId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	// The author of the room cannot ban himself
	if blockingUserId < 1 || blockingUserId == authorData.UserId {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	strStatus := params("status")
	status, err := strconv.ParseBool(strStatus)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = hd.RoomController.Ban(blockingUserId, roomId, status)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
