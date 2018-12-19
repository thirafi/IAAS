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

func GetAllCompany(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	company := []model.Company{}
	db.Find(&company)
	respondJSON(w, http.StatusOK, company)
}

func CreateCompany(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	company := model.Company{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&company); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&company).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, company)
}

func GetCompany(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	company := getCompanyOr404(db, id, w, r)
	if company == nil {
		return
	}
	respondJSON(w, http.StatusOK, company)
}

// func UpdateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	id := vars["id"]
// 	user := getUserOr404(db, id, w, r)
// 	if user == nil {
// 		return
// 	}

// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&user); err != nil {
// 		respondError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	defer r.Body.Close()

// 	if err := db.Save(&user).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusOK, user)
// }

// func DeleteUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	id := vars["id"]
// 	user := getUserOr404(db, id, w, r)
// 	if user == nil {
// 		return
// 	}
// 	if err := db.Delete(&user).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusNoContent, nil)
// }

// func DisableUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	id := vars["id"]
// 	user := getUserOr404(db, id, w, r)
// 	if user == nil {
// 		return
// 	}
// 	user.Disable()
// 	if err := db.Save(&user).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusOK, user)
// }

// func EnableUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	id := vars["id"]
// 	user := getUserOr404(db, id, w, r)
// 	if user == nil {
// 		return
// 	}
// 	user.Enable()
// 	if err := db.Save(&user).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusOK, user)
// }

// getUserOr404 gets a user instance if exists, or respond the 404 error otherwise
func getCompanyOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Company {
	company := model.Company{}
	ud, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	ad := uint(ud)
	if err := db.First(&company, ad).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &company
}
