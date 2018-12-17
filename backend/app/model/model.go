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
	Health       Health `json:"health"`
	Address      Address `json:"address"`
	Contact			 Contact `json:"contact"`
	Language	   []Language `json:"language"`
	Period			 []Period `json:"period"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

type Health struct {
	Id_user      uint
	Smoke        bool `json:"smoke"`
	Problem      string `json:"problem"`
	Restriction  string `json:restriction`
}

type Address struct {
	Id_user      uint
	City         string `json:"city"`
	Province     string `json:"province"`
	PostalCode   uint `json:"postalcode"`
}

type Contact struct {
	Id_user      uint
	Phone				uint `json:"phone"`
	EmergencyPhone 		uint `json:"emergencyphone"`
	EmergencyName  		string `json:"emergencyname"`
	EmergencyRelation string `json:"emergencyrelation"`
}

type Language struct {
	Id_user      uint
	Language	  string `json:"language"`
	Speaking	  string `json:"speaking"`
	Structure		string `json:"structure"`
	Reading			string `json:"reading"`
	Writing			string `json:"writing"`
}

type Period struct {
	Id_user      uint
	Start       string `json:"start"`
	Finish      string `json:"finish"`
	Length		  string `json:"length"`
}

type Company struct {
	Id_company    uint   `gorm:"primary_key"`
	Com_Name      string ` json:"comname"`
	Com_Desc      string ` json:"comdesc"`
	Com_Image     string ` json:"comimage"`
	Com_Portfolio string ` json:"portfolio"`
	Com_Slogan    string ` json:"slogan"`
	Com_Address   string ` json:"address"`
	Com_Website   string `json:"website"`
	Com_Facebook  string `json:"facebook"`
	Com_Linkedin  string `json:"linkedin"`
	Com_Instagram string `json:"instagram"`
	Com_Google    string `json:"google"`
}

type Job struct {
	Id_job       uint   `gorm:"primary_key"`
	Job_Title    string ` json:"jobtitle"`
	Job_Desc     string ` json:"jobdesc"`
	Job_Address  string ` json:"jobaddress"`
	Job_Image    string ` json:"jobimage"`
	Job_Category string ` json:"cat"`
	Job_Type     string ` json:"type"`
	Job_Location string `json:"location"`
	Job_Tag      string `json:"tag"`
	// Job_ClosingDate    string  `json:"closingdate"`
	Job_NotifEmail string `json:"notifemail"`
	Job_MainAct    string `json:"mainact"`
	ClosedAt       time.Time
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

func (e *Job) Disable() {
	e.Status = false
}

func (p *Job) Enable() {
	p.Status = true
}

// db.Model(&user).Related(&health)
// db.Model(&user).Related(&address)
// db.Model(&user).Related(&contact)
// db.Model(&user).Related(&languages)
// db.Model(&user).Related(&periods)

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Employee{}, &User{}, &Health{}, &Address{}, &Contact{}, &Language{}, &Period{})
	return db
}
