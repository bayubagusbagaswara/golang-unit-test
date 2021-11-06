package helper

import "testing"

func TestHelloWorld(t *testing.T) {

	// simpan hasil dari function yang kita test, misal simpan ke result
	result := HelloWorld("Bayu")
	// lalu cek isi dari resultnya apakah sesuai dengan ekspektasi atau tidak, kalau tidak sesuai, maka lempar error

	if result != "Hello Bayu" {
		// jika tidak sama dengan Hello Bayu, maka error
		panic("Result is not 'Hello Bayu'")
	}

}
