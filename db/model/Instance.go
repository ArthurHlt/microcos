package model

import (
	"github.com/jinzhu/gorm"
)
type Instance struct {
	gorm.Model
	InstanceName string        `sql:"type:varchar(100);unique_index"`
	Applications []Application
}