package model

import (
	"github.com/jinzhu/gorm"
)
type Application struct {
	gorm.Model
	AppName               string `sql:"type:varchar(100);unique_index"`
	HostName              string `sql:"type:varchar(100);unique_index"`
	StatusPageUrl         string `sql:"not null"`
	Port                  int    `sql:"not null"`
	SecurePort            bool   `sql:"not null"`
	Threshold             int    `sql:"not null"`
	RenewalIntervalInSecs int    `sql:"not null"`
	Proxy                 string
	InstanceID            int    `sql:"index"`
	ExpectedStatusCode    int    `sql:"not null"`
}