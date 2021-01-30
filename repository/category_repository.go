package repository

import "latihan-belajar-golang-unit-test/entity"

type CategoryRepository interface { // ini merupakan kontrak
	// memiliki function FindById dengan parameter id string serta return valuenya
	// pointer dari package entity struct category
	FindById(id string) *entity.Category // jadi returnnya apabila ada di database maka category, jika tidak ada maka nil
}
