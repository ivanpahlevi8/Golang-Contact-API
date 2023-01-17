package main

import (
	"myapi/pkg/handler"
	"myapi/pkg/model"
	"myapi/pkg/repository"
	"myapi/pkg/service"
	"net/http"
)

func main() {

	// init repo
	var myRepo []model.Account
	repoInterface := repository.InitRepository(myRepo)

	// init service
	serviceInterface := service.InitService(repoInterface)

	// init handler
	handlerInterface := handler.InitHandler(serviceInterface)

	// start routing
	http.HandleFunc("/add", handlerInterface.AddData)
	http.HandleFunc("/getusername", handlerInterface.GetDataByUsername)
	http.HandleFunc("/getpassword", handlerInterface.GetDataByPassword)
	http.HandleFunc("/updatepassword", handlerInterface.UpdateDataByPassword)
	http.HandleFunc("/updateusername", handlerInterface.UpdateDataByUsername)
	http.HandleFunc("/deleteusername", handlerInterface.DeleteDataByUsername)
	http.HandleFunc("/deletepassword", handlerInterface.DeleteDataByPassword)

	http.ListenAndServe(":3030", nil)
}
