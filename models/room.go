package models

type Room struct {
	RoomId    int    `json:"room_id"`
	Name      string `json:"name"`
	UUID      string `json:"uuid"`
	SubjectId int    `json:"subject_id"`
	AuthorId  int    `json:"author_id"`
}
