package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Flyewzz/group_preparation/auth"
	"github.com/gorilla/mux"
)

func (hd *HandlerData) GetRoomHandler(w http.ResponseWriter, r *http.Request) {
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
}

func (hd *HandlerData) BanRoomHandler(w http.ResponseWriter, r *http.Request) {
}
