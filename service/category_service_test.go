package service

import (
	"golang-unit-test/entity"
	"golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// bikin unit test dari Mock

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}

// kita set dari categoryRepository
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_Get(t *testing.T) {

	// program mock
	// ketikan FindById dipanggil dan parameternya adalah 1
	// maka kita return datanya, misalnya adalah nil atau datanya kosong
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	// jalankan function categoryService
	category, err := categoryService.Get("1")

	assert.Nil(t, category)
	assert.NotNil(t, err)

}

// function test jika berhasil
func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "Laptop",
	}
	// panggil object Mock
	categoryRepository.Mock.On("FindById", "2").Return(category)

	// jika kita memanggil 2, maka balikannya adalah entity diatas
	result, err := categoryService.Get("2")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
