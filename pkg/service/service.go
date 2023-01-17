package service

import (
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

func (m *Service) AddData(data model.Account) {
	m.repo.AddData(data)
}

func (m *Service) GetDataByUsername(username string) model.Account {
	var allData = m.repo.GetAllData()
	var getData model.Account
	for _, data := range allData {
		if data.GetUsername() == username {
			getData = data
			break
		}
	}

	return getData
}

func (m *Service) GetDataByPassword(password string) model.Account {
	var allData = m.repo.GetAllData()
	var getData model.Account

	for _, data := range allData {
		if data.GetPassword() == password {
			getData = data
			break
		}
	}

	return getData
}

func (m *Service) UpdateDataByUsername(username string, data model.Account) model.Account {
	var allData = m.repo.GetAllData()
	var i = 0

	for _, data := range allData {
		if data.GetUsername() == username {
			break
		}
		i++
	}

	m.repo.UpdateData(i, data)

	return myService.repo.GetData(i)
}

func (m *Service) UpdateDataByPassword(password string, data model.Account) model.Account {
	var allData = m.repo.GetAllData()
	var i = 0

	for _, data := range allData {
		if data.GetPassword() == password {
			break
		}
		i++
	}

	m.repo.UpdateData(i, data)

	return myService.repo.GetData(i)
}

func (m *Service) DeleteDataByUsername(username string) model.Account {
	var allData = m.repo.GetAllData()
	var i = 0

	for _, data := range allData {
		if data.GetUsername() == username {
			break
		}
		i++
	}

	getData := m.repo.GetData(i)
	m.repo.DeleteData(i)

	return getData
}

func (m *Service) DeleteDataByPassword(password string) model.Account {
	var allData = m.repo.GetAllData()
	var i = 0

	for _, data := range allData {
		if data.GetPassword() == password {
			break
		}
		i++
	}

	getData := m.repo.GetData(i)
	m.repo.DeleteData(i)

	return getData
}
