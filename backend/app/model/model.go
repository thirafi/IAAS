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

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Employee{}, &User{})
	return db
}
