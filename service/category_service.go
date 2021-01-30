package service

import (
	"errors"
	"latihan-belajar-golang-unit-test/entity"
	"latihan-belajar-golang-unit-test/repository"
)

type CategoryService struct {
	// bayangkan kalau ini merupakan third party database
	// NOTE: Unit Test itu idealnya tidak running third party apapun, jadi tidak running
	//database, web server dan lainnya.
	Repository repository.CategoryRepository
	// oleh karena itu kita akan membuat mock object
}

// membuat function Get dengan parameter id string dan return value category atau error
// jika ada maka return value pointer category
// jika tidak ada maka menampilkan error
func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id) // ini mencari Repository dari Func FindById sesuai dengan parameter func Get
	if category == nil {
		// jika nil maka, category menjadi nil dan akan menampilkan errors
		return nil, errors.New("Category Not Found")
	} else {
		// jika category ada maka, akan menampilkan category tersebut dan error tersebut nil
		return category, nil
	}
}
