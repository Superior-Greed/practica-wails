package models

import "gorm.io/gorm"

type Folder struct {
	gorm.Model
	Name             string `json:"name" db:"name"`
	Description      string `json:"description" db:"description"`
	Rute             string `json:"rute" db:"rute"`
	TerminationImage string `json:"termination_image" db:"termination_image"`
}
