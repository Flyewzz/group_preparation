package room

type RoomData struct {
	RoomId      int    `json:"room_id"`
	Name        string `json:"name"`
	UUID        string `json:"uuid"`
	SubjectId   string `json:"subject_id,omitempty"`
	AuthorEmail string `json:"author_email"`
	Type        string `json:"type"`
	Banned      bool   `json:"banned"`
}
