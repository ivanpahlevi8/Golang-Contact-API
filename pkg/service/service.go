package service

import (
	"errors"
	"log"
	"myapi/pkg/model"
	"myapi/pkg/repository"
)

var myService Service

type Service struct {
	repo repository.Repository
}

func InitService(repo repository.Repository) Service {
	myService.repo = repo
	return myService
}

func (m *Service) AddData(data model.Account) (model.Account, error) {
	// check input data
	errorCheck := false

	if data.GetUsername() == "" || data.GetUsername() == " " || data.GetUsername() == "  " || data.GetUsername() == "   " || data.GetUsername() == "    " || data.GetUsername() == "     " {
		errorCheck = true
	} else if data.GetPassword() == "" || data.GetPassword() == " " || data.GetPassword() == "  " || data.GetPassword() == "   " || data.GetPassword() == "    " || data.GetPassword() == "     " {
		errorCheck = true
	}

	log.Println("Input pass :", data.GetPassword())

	if errorCheck {
		return data, errors.New("Invalid Input Error")
	} else {
		m.repo.AddData(data)
		return data, nil
	}
}

func (m *Service) GetDataByUsername(username string) (model.Account, error) {
	var allData = m.repo.GetAllData()
	var getData model.Account
	i := 0
	for _, data := range allData {
		if data.GetUsername() == username {
			getData = data
			break
		}
		i++
	}

	if i == len(allData) {
		return getData, errors.New("Username not found")
	} else {
		return getData, nil
	}
}

func (m *Service) GetDataByPassword(password string) (model.Account, error) {
	var allData = m.repo.GetAllData()
	var getData model.Account
	i := 0
	for _, data := range allData {
		if data.GetPassword() == password {
			getData = data
			break
		}
		i++
	}

	if i == len(allData) {
		return getData, errors.New("Password Not Found")
	} else {
		return getData, nil
	}
}

func (m *Service) UpdateDataByUsername(username string, data model.Account) (model.Account, error) {
	var allData = m.repo.GetAllData()
	var i = 0

	for _, data := range allData {
		if data.GetUsername() == username {
			break
		}
		i++
	}

	if i == len(allData) {
		m.repo.UpdateData(i, data)
		return myService.repo.GetData(i), nil
	} else {
		return model.Account{}, errors.New("Cannot Update By Username")
	}

}

func (m *Service) UpdateDataByPassword(password string, data model.Account) (model.Account, error) {
	var allData = m.repo.GetAllData()
	var i = 0

	for _, data := range allData {
		if data.GetPassword() == password {
			break
		}
		i++
	}

	if i == len(allData) {
		m.repo.UpdateData(i, data)
		return myService.repo.GetData(i), nil
	} else {
		return model.Account{}, errors.New("cannot update by username")
	}
}

func (m *Service) DeleteDataByUsername(username string) (model.Account, error) {
	var allData = m.repo.GetAllData()
	var i = 0

	for _, data := range allData {
		if data.GetUsername() == username {
			break
		}
		i++
	}

	if i == len(allData) {
		getData := m.repo.GetData(i)
		m.repo.DeleteData(i)
		return getData, nil
	} else {
		return model.Account{}, errors.New("error when delete data with username")
	}
}

func (m *Service) DeleteDataByPassword(password string) (model.Account, error) {
	var allData = m.repo.GetAllData()
	var i = 0

	for _, data := range allData {
		if data.GetPassword() == password {
			break
		}
		i++
	}

	if i == len(allData) {
		getData := m.repo.GetData(i)
		m.repo.DeleteData(i)
		return getData, nil
	} else {
		return model.Account{}, errors.New("error when delete data with username")
	}
}
