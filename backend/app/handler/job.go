package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllJob(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	jobs := []model.Job{}
	db.Find(&jobs)
	respondJSON(w, http.StatusOK, jobs)
}

func CreateJob(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	job := model.Job{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&job); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&job).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, job)
}

func GetJob(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	job := getJobOr404(db, id, w, r)
	if job == nil {
		return
	}
	respondJSON(w, http.StatusOK, job)
}

// getUserOr404 gets a user instance if exists, or respond the 404 error otherwise
func getJobOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Job {
	job := model.Job{}
	ud, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	ad := uint(ud)
	if err := db.First(&job, ad).Error; err != nil {
		fmt.Println("udah jalan coi:", job)
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &job
}

// func getjobOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Job {
// 	job := model.Job{}
// 	ud, err := strconv.ParseUint(id, 10, 64)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	ad := uint(ud)
// 	if err := db.First(&job, model.Job{ID: ad}).Error; err != nil {
// 		respondError(w, http.StatusNotFound, err.Error())
// 		return nil
// 	}
// 	return &job
// }
