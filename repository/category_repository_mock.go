package repository

import (
	"github.com/stretchr/testify/mock"
	"latihan-belajar-golang-unit-test/entity"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock // punyanya testify
}

// implementasi untuk struct CategoryRepositoryMock

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category { // kembaliannya hanya satu data yaitu Category
	// memanggil parameter Id maka kita panggil Mock nya
	arguments := repository.Mock.Called(id)

	// lalu cek
	if arguments.Get(0) == nil {
		// kalau returnnya nil maka nil
		return nil
	} else {
		// jika tidak maka akan mengembalikan categorynya
		// melakukan konversi interface ke dalam struct
		category := arguments.Get(0).(entity.Category) // karena ini pointer tanda *
		return &category                               // maka tambahkan tanda &
	}
}
