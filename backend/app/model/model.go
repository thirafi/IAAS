package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Employee struct {
	gorm.Model
	Name   string `gorm:"unique" json:"name"`
	City   string `json:"city"`
	Age    int    `json:"age"`
	Status bool   `json:"status"`
}
type User struct {
	Id_user      uint   `gorm:"primary_key"`
	FirstName    string ` json:"firstname"`
	LastName     string ` json:"lLastname"`
	Gender       string ` json:"gender"`
	DateofBirth  string ` json:"dateofbirth"`
	PlaceofBirth string ` json:"placeofbirth"`
	Hometown     string ` json:"hometown"`
	Role         string `json:"city"`
	Nationality  int    `json:"age"`
	Status       bool   `json:"status"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

type Company struct{
	Id_company   	uint   `gorm:"primary_key"`
	Com_Name    	string ` json:"comname"`
	Com_Desc     	string ` json:"comdesc"`
	Com_Image       string ` json:"comimage"`
	Com_Portfolio  	string ` json:"portfolio"`
	Com_Slogan 		string ` json:"slogan"`
	Com_Address     string ` json:"address"`
	Com_Website     string 	`json:"website"`
	Com_Facebook  	string  `json:"facebook"`
	Com_Linkedin    string  `json:"linkedin"`
	Com_Instagram	string  `json:"instagram"`
	Com_Google		string  `json:"google"`
}

func (e *Employee) Disable() {
	e.Status = false
}

func (p *Employee) Enable() {
	p.Status = true
}

func (e *User) Disable() {
	e.Status = false
}

func (p *User) Enable() {
	p.Status = true
}

func (e *Company) Disable() {
	e.Status = false
}

func (p *Company) Enable() {
	p.Status = true
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Employee{}, &User{})
	return db
}
