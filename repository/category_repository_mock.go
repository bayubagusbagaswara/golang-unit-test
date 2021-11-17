package repository

import (
	"golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

// bikin struct
type CategoryRepositoryMock struct {
	Mock mock.Mock
}

// bikin function yang dijadikan Mock
func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {

	// bikin program untuk si mock nya
	// panggil Mocknya dan masukkan parameter sesuai dengan parameter yang ada di function yang dibuatkan mock
	// balikannya adalah argument
	arguments := repository.Mock.Called(id)

	// karena balikannya bisa nil, maka kita harus cek dulu
	if arguments.Get(0) == nil {
		return nil
	} else {
		// jika return nya tidak sama dengan nil, maka kita balikkan data category nya
		category := arguments.Get(0).(entity.Category)
		// karena category ini adalah pointer, maka kita tambahkan &
		return &category
	}

}
