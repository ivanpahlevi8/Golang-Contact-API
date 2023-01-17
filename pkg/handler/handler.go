package handler

import (
	"encoding/json"
	"log"
	"myapi/pkg/model"
	"myapi/pkg/service"
	"net/http"
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

	m.myService.AddData(account)

	log.Println("Success Adding Data")
}

func (m *Handler) GetDataByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := r.URL.Query()

	getUsername := params.Get("username")

	getAccount := m.myService.GetDataByUsername(getUsername)

	json.NewEncoder(w).Encode(getAccount)

	log.Println("Success get Data By Username")
}

func (m *Handler) GetDataByPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := r.URL.Query()

	getPassword := params.Get("password")

	getAccount := m.myService.GetDataByPassword(getPassword)

	json.NewEncoder(w).Encode(getAccount)

	log.Println("Success get Data By Username")
}

func (m *Handler) UpdateDataByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var account model.Account
	_ = json.NewDecoder(r.Body).Decode(&account)

	params := r.URL.Query()
	getUsername := params.Get("username")

	getAccount := m.myService.UpdateDataByUsername(getUsername, account)

	json.NewEncoder(w).Encode(getAccount)

	log.Println("Success update data by username")
}

func (m *Handler) UpdateDataByPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var account model.Account
	_ = json.NewDecoder(r.Body).Decode(&account)

	params := r.URL.Query()
	getPassword := params.Get("password")

	getAccount := m.myService.UpdateDataByPassword(getPassword, account)

	json.NewEncoder(w).Encode(getAccount)

	log.Println("Success update data by password")
}

func (m *Handler) DeleteDataByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := r.URL.Query()
	getUsername := params.Get("username")

	getAccount := m.myService.DeleteDataByUsername(getUsername)

	json.NewEncoder(w).Encode(getAccount)

	log.Println("Success Delete Data By Username")
}

func (m *Handler) DeleteDataByPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := r.URL.Query()
	getPassword := params.Get("password")

	getAccount := m.myService.DeleteDataByPassword(getPassword)

	json.NewEncoder(w).Encode(getAccount)

	log.Println("Success Delete Data By Password")
}
