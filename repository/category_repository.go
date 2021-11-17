package repository

import "golang-unit-test/entity"

// bikin repository
// tapi jangan langsung bikin function ke databasenya
// tapi bikin interface nya dulu, lalu bikin kontrak functionnya
type CategoryRepository interface {

	// dimana balikannya adalah pointer entity Category
	FindById(id string) *entity.Category
}
