package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Cliente struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Nome     string `json:"nome" validate:"nonzero"`
	telefone string `json:"telefone" validate:"len=11, regexp=^[0-9]*$"`
}

func validateCliente(cliente *Cliente) error {
	if err := validator.Validate(cliente); err != nil {
		return err
	}
	return nil
}
