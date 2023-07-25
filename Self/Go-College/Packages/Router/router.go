package router

import (
	controller "github.com/SudebDolui/CollegeApi/Packages/Controller"
	"github.com/gorilla/mux"
)

var RegisterRouter = func(router *mux.Router) {
	router.HandleFunc("/api/registerstudent", controller.RegisterStudent).Methods("POST")
	router.HandleFunc("/api/getstudents", controller.GetStudents).Methods("GET")
	router.HandleFunc("/api/getstudent/{rollnumber}", controller.GetStudent).Methods("GET")
	// router.HandleFunc("/api/updatestudent/{rollnumber}").Methods("PUT")
	// router.HandleFunc("/api/deletestudent/{rollnumber}").Methods("DELETE")
	// router.HandleFunc("/api/deletestudents").Methods("DELETE")

}
