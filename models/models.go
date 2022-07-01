package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name     string `json:"name" validate:"nonzero"`
	Document string `json:"document" validate:"len=11,regexp=^[0-9]*$"`
	Age      int    `json:"age" validate:"nonzero"`
}

func ValidateStudents(student *Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}
	return nil
}
