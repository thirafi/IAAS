package model

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//a struct to rep user account
type Token struct {
	Id_user uint
	jwt.StandardClaims
}
type Account struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
	Role     string `json:"city"`
	
}

type User struct {
	gorm.Model
	Account      Account
	AccountID    uint
	FirstName    string ` json:"firstname"`
	LastName     string ` json:"lLastname"`
	Gender       string ` json:"gender"`
	DateofBirth  string ` json:"dateofbirth"`
	PlaceofBirth string ` json:"placeofbirth"`
	Hometown     string ` json:"hometown"`
	Nationality  int    `json:"nationality"`
	Status       bool   `json:"status"`
	// Language []Language
	Health       Health     `json:"health"`
	Address      Address    `json:"address"`
	Contact      Contact    `json:"contact"`
	Language     []Language `json:"language"`
	Period       []Period   `json:"period"`
}

type Health struct {
	gorm.Model
	// Account     Account
	AccountID   uint
	Smoke       bool
	Problem     string
	Restriction string
}

type Address struct {
	gorm.Model
	// Account    Account
	AccountID  uint
	City       string `json:"city"`
	Province   string `json:"province"`
	PostalCode uint   `json:"postalcode"`
}

type Contact struct {
	gorm.Model
	// Account           Account
	AccountID         uint
	Phone             uint   `json:"phone"`
	EmergencyPhone    uint   `json:"emergencyphone"`
	EmergencyName     string `json:"emergencyname"`
	EmergencyRelation string `json:"emergencyrelation"`
}

type Language struct {
	gorm.Model
	Language  string `json:"language"`
	Speaking  string `json:"speaking"`
	Structure string `json:"structure"`
	Reading   string `json:"reading"`
	Writing   string `json:"writing"`
	AccountID uint
}

type Period struct {
	gorm.Model
	// Account   Account
	AccountID uint
	Start     string `json:"start"`
	Finish    string `json:"finish"`
	Length    string `json:"length"`
}

type Company struct {
	gorm.Model
	Name      string
	Desc      string
	Logo      string
	Cover     string
	Portfolio string
	Slogan    string
	Address   string
	Website   string
	Facebook  string
	Linkedin  string
	Instagram string
	Google    string
}

type Job struct {
	gorm.Model
	Company    Company
	CompanyID  uint
	Title      string
	Desc       string
	Address    string
	Image      string
	Category   string
	Type       string
	Location   string
	Tag        string // Job_ClosingDate    string  `json:"closingdate"`
	NotifEmail string
	Activity   string
	ClosedAt   time.Time
}

type Prerequisite struct {
	gorm.Model
	Job           Job
	JobId         uint
	Preference    string
	Nationality   string
	Language      string
	StudyLevel    string
	SelectProcess string
	Logistic      []Logistic
	AddSelection  []AddSelection
	Working       *time.Time
	Duration      *time.Time
	PeriodAvaible *time.Time
}
type AddSelection struct {
	gorm.Model
	PrerequisiteID uint
	Additional     string
}

type Logistic struct {
	gorm.Model
	PrerequisiteID uint
	Type           string
}

type Employee struct {
	gorm.Model
	Name   string `gorm:"unique" json:"name"`
	City   string `json:"city"`
	Age    int    `json:"age"`
	Status bool   `json:"status"`
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
	db.AutoMigrate(&Employee{}, &User{}, &Health{}, &Address{}, &Contact{}, &Language{}, &Period{}, &Account{}, &Company{}, &Job{}, &Prerequisite{}, &AddSelection{}, &Logistic{})
	return db
}

