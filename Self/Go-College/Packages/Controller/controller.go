package controller

import (
	"encoding/json"
	"net/http"

	config "github.com/SudebDolui/CollegeApi/Packages/Config"
	models "github.com/SudebDolui/CollegeApi/Packages/Models"
	utils "github.com/SudebDolui/CollegeApi/Packages/Utils"
	"github.com/gorilla/mux"
)

func RegisterStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var student models.Students
	_ = json.NewDecoder(r.Body).Decode(&student)
	config.HelperRegisterStudent(student)
	json.NewEncoder(w).Encode(student)
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	students := config.HelperGetStudents()
	json.NewEncoder(w).Encode(students)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	params := mux.Vars(r)
	student := config.HelperGetStudent(params["rollnumber"])
	json.NewEncoder(w).Encode(student)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	var student *models.Students{}
	utils.ParseBody(r, &student)
	student := config.HelperUpdateStudent(&student, params["rollnumber"])
	json.NewEncoder(w).Encode(student)
}
