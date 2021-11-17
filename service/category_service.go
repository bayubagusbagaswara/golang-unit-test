package service

import (
	"errors"
	"golang-unit-test/entity"
	"golang-unit-test/repository"
)

// bikin struct
type CategoryService struct {
	Repository repository.CategoryRepository
}

// kita bikin method atau function untuk service
// balikannya adalah jika ada datanya, maka balikan entity Category
// jika tidak ada datanya maka balikkan error
func (service CategoryService) Get(id string) (*entity.Category, error) {
	// panggil Repository
	category := service.Repository.FindById(id)

	// cek
	if category == nil {
		// balikannya adalah datanya nil, dengan keterangannya adalah erros
		return nil, errors.New("Category Not Found")
	} else {
		// balikanny adalah datanya berisi category, dan errornya adalah nil (tidak terjadi error)
		return category, nil
	}

}
