package repository

import "myapi/pkg/model"

var repo Repository

type Repository struct {
	account []model.Account
}

func InitRepository(input []model.Account) Repository {
	repo.account = input
	return repo
}

func (m *Repository) AddData(data model.Account) {
	repo.account = append(repo.account, data)
}

func (m *Repository) GetData(index int) model.Account {
	data := repo.account[index]

	return data
}

func (m *Repository) UpdateData(index int, data model.Account) {
	repo.account[index] = data
}

func (m *Repository) DeleteData(index int) {
	repo.account = append(repo.account[:index], repo.account[index+1:]...)
}

func (m *Repository) GetAllData() []model.Account {
	return repo.account
}
