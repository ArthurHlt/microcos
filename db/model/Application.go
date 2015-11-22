package model

import (
	"github.com/jinzhu/gorm"
)
type Application struct {
	gorm.Model
	AppName               string `sql:"not null"`
	HostName              string `sql:"type:varchar(100);unique_index"`
	StatusPageUrl         string `sql:"not null"`
	Port                  int    `sql:"not null"`
	SecurePort            bool   `sql:"not null"`
	Threshold             int    `sql:"not null"`
	RenewalIntervalInSecs int    `sql:"not null"`
	InstanceID            int    `sql:"index"`
}