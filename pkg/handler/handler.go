package handler

import (
	"encoding/json"
	"log"
	"myapi/pkg/model"
	"myapi/pkg/service"
	"net/http"
	"time"
)

var MyHandler Handler

type Handler struct {
	myService service.Service
}

func InitHandler(service service.Service) Handler {
	MyHandler.myService = service
	return MyHandler
}

func (m *Handler) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatiom/json")
	var account model.Account

	_ = json.NewDecoder(r.Body).Decode(&account)

	data, err := m.myService.AddData(account)

	if err != nil {
		t := time.Now().Format("2006-01-02 15:04:05")
		errRespone := model.Error{
			ErrorMessage: err.Error(),
			TimeStamp:    t,
		}
		json.NewEncoder(w).Encode(errRespone)
		w.WriteHeader(404)
		log.Println("Error Adding Data")
	} else {
		json.NewEncoder(w).Encode(data)
		log.Println("Success Adding Data")
		w.WriteHeader(201)
	}
}

func (m *Handler) GetDataByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := r.URL.Query()

	getUsername := params.Get("username")

	getAccount, err := m.myService.GetDataByUsername(getUsername)

	if err != nil {
		t := time.Now()
		errRespone := model.Error{
			ErrorMessage: err.Error(),
			TimeStamp:    t.String(),
		}
		json.NewEncoder(w).Encode(errRespone)
		log.Println("Error get data by username")
	} else {
		json.NewEncoder(w).Encode(getAccount)
		log.Println("Success get Data By Username")
	}
}

func (m *Handler) GetDataByPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := r.URL.Query()

	getPassword := params.Get("password")

	getAccount, err := m.myService.GetDataByPassword(getPassword)

	if err != nil {
		t := time.Now()
		errRespone := model.Error{
			ErrorMessage: err.Error(),
			TimeStamp:    t.String(),
		}
		json.NewEncoder(w).Encode(errRespone)
		log.Println("Error get data by password")
	} else {
		json.NewEncoder(w).Encode(getAccount)
		log.Println("Success get Data By password")
	}
}

func (m *Handler) UpdateDataByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var account model.Account
	_ = json.NewDecoder(r.Body).Decode(&account)

	params := r.URL.Query()
	getUsername := params.Get("username")

	getAccount, err := m.myService.UpdateDataByUsername(getUsername, account)

	if err != nil {
		t := time.Now()
		errRespone := model.Error{
			ErrorMessage: err.Error(),
			TimeStamp:    t.String(),
		}
		json.NewEncoder(w).Encode(errRespone)
		log.Println("Error update data by username")
	} else {
		json.NewEncoder(w).Encode(getAccount)
		log.Println("Success update data by username")
	}
}

func (m *Handler) UpdateDataByPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var account model.Account
	_ = json.NewDecoder(r.Body).Decode(&account)

	params := r.URL.Query()
	getPassword := params.Get("password")

	getAccount, err := m.myService.UpdateDataByPassword(getPassword, account)

	if err != nil {
		t := time.Now()
		errRespone := model.Error{
			ErrorMessage: err.Error(),
			TimeStamp:    t.String(),
		}
		json.NewEncoder(w).Encode(errRespone)
		log.Println("Error update data by password")
	} else {
		json.NewEncoder(w).Encode(getAccount)
		log.Println("Success update data by password")
	}
}

func (m *Handler) DeleteDataByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := r.URL.Query()
	getUsername := params.Get("username")

	getAccount, err := m.myService.DeleteDataByUsername(getUsername)

	if err != nil {
		t := time.Now()
		errRespone := model.Error{
			ErrorMessage: err.Error(),
			TimeStamp:    t.String(),
		}
		json.NewEncoder(w).Encode(errRespone)
		log.Println("Error delete data by username")
	} else {
		json.NewEncoder(w).Encode(getAccount)
		log.Println("Success delete data by username")
	}
}

func (m *Handler) DeleteDataByPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := r.URL.Query()
	getPassword := params.Get("password")

	getAccount, err := m.myService.DeleteDataByPassword(getPassword)

	if err != nil {
		t := time.Now()
		errRespone := model.Error{
			ErrorMessage: err.Error(),
			TimeStamp:    t.String(),
		}
		json.NewEncoder(w).Encode(errRespone)
		log.Println("Error delete data by password")
	} else {
		json.NewEncoder(w).Encode(getAccount)
		log.Println("Success delete data by password")
	}

}
