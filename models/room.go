package models

type Room struct {
	RoomId    int    `json:"room_id"`
	Name      string `json:"name"`
	UUID      string `json:"uuid"`
	TypeId    int    `json:"type_id"`
	SubjectId int    `json:"subject_id"`
	AuthorId  int    `json:"author_id"`
}
