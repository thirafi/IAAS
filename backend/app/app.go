package app

import (
	"fmt"
	"log"
	"net/http"

	"./config"
	"./handler"
	"./model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling employee
	a.Get("/", a.HelloWorld)
	a.Get("/employees", a.GetAllEmployees)
	a.Post("/employees", a.CreateEmployee)
	a.Get("/employees/{name}", a.GetEmployee)
	a.Put("/employees/{name}", a.UpdateEmployee)
	a.Delete("/employees/{name}", a.DeleteEmployee)
	a.Put("/employees/{name}/disable", a.DisableEmployee)
	a.Put("/employees/{name}/enable", a.EnableEmployee)
	// Routing for handling user
	a.Get("/users", a.GetAllUsers)
	a.Post("/users", a.CreateUser)
	a.Get("/users/{id}", a.GetUser)
	a.Put("/users/{id}", a.UpdateUser)
	a.Delete("/users/{id}", a.DeleteUser)
	a.Put("/users/{id}/disable", a.DisableUser)
	a.Put("/users/{id}/enable", a.EnableUser)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa ini dia ")
}

// Handlers to manage Employee Data
func (a *App) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	handler.GetAllEmployees(a.DB, w, r)
}

func (a *App) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	handler.CreateEmployee(a.DB, w, r)
}

func (a *App) GetEmployee(w http.ResponseWriter, r *http.Request) {
	handler.GetEmployee(a.DB, w, r)
}

func (a *App) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	handler.UpdateEmployee(a.DB, w, r)
}

func (a *App) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	handler.DeleteEmployee(a.DB, w, r)
}

func (a *App) DisableEmployee(w http.ResponseWriter, r *http.Request) {
	handler.DisableEmployee(a.DB, w, r)
}

func (a *App) EnableEmployee(w http.ResponseWriter, r *http.Request) {
	handler.EnableEmployee(a.DB, w, r)
}

// Handlers to manage User Data
func (a *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	handler.GetAllUsers(a.DB, w, r)
}

func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	handler.CreateUser(a.DB, w, r)
}

func (a *App) GetUser(w http.ResponseWriter, r *http.Request) {
	handler.GetUser(a.DB, w, r)
}

func (a *App) UpdateUser(w http.ResponseWriter, r *http.Request) {
	handler.UpdateUser(a.DB, w, r)
}

func (a *App) DeleteUser(w http.ResponseWriter, r *http.Request) {
	handler.DeleteUser(a.DB, w, r)
}

func (a *App) DisableUser(w http.ResponseWriter, r *http.Request) {
	handler.DisableUser(a.DB, w, r)
}

func (a *App) EnableUser(w http.ResponseWriter, r *http.Request) {
	handler.EnableUser(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
