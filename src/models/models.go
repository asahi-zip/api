package models

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
	Token     string `gorm:"unique"`
	OrgID     *uint
	Org       Org `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Org struct {
	gorm.Model
	Name   string
	Users  []User `gorm:"foreignKey:OrgID"`
	Medias []Media
}

type Media struct {
	gorm.Model
	FileName string
	FileData []byte
	FileSize int64
	MimeType string
	OrgID    uint
	Org      Org
}
