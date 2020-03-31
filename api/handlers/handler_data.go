package handlers

import (
	. "github.com/Flyewzz/group_preparation/interfaces"
)

type HandlerData struct {
	UniversityController UniversityController
	SubjectController    SubjectController
	MaterialController   MaterialController
	AuthController       AuthController
}

func NewHandlerData(
	uc UniversityController,
	sc SubjectController,
	mc MaterialController,
	ac AuthController,
) *HandlerData {
	return &HandlerData{
		UniversityController: uc,
		SubjectController:    sc,
		MaterialController:   mc,
		AuthController:       ac,
	}
}
