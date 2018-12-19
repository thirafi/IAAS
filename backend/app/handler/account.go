package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"../model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Register(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	account := model.Account{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&account); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)
	if err := db.Save(&account).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	tk := model.Token{Id_user: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString
	account.Password = "" //delete password
	respondJSON(w, http.StatusCreated, account)
}
func Login(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	account := model.Account{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&account); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	account1 := geAccountOr404(db, account.Email, w, r)

	arr := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(account1.Password))
	if arr != nil && arr == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		respondError(w, http.StatusNotFound, arr.Error())
	}
	//Worked! Logged In
	account.Password = ""

	//Create JWT token
	tk := model.Token{Id_user: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString
	respondJSON(w, http.StatusCreated, account)
}

func geAccountOr404(db *gorm.DB, email string, w http.ResponseWriter, r *http.Request) *model.Account {
	account := model.Account{}
	if err := db.Where("email = ?", email).First(&account).Error; err != nil {
		fmt.Println("akun id ", account.ID)
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &account
}
