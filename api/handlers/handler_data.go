package handlers

import (
	. "github.com/Flyewzz/group_preparation/interfaces"
)

type HandlerData struct {
	UniversityController UniversityController
	SubjectController    SubjectController
	MaterialController   MaterialController
	AuthController       AuthController
	RoomController       RoomController
}

func NewHandlerData(
	uc UniversityController,
	sc SubjectController,
	mc MaterialController,
	ac AuthController,
	rc RoomController,
) *HandlerData {
	return &HandlerData{
		UniversityController: uc,
		SubjectController:    sc,
		MaterialController:   mc,
		AuthController:       ac,
		RoomController:       rc,
	}
}
