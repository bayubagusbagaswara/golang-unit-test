package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Bayu")

	if result != "Hello Bayu" {
		// error
		t.Fail()
	}
	// kode program ini tetap dieksekusi, tapi hasil dari unit testnya tetap Fail
	fmt.Println("Test HelloWorld Tetap Dieksekusi Selesai")
}

func TestHelloWorldOther(t *testing.T) {
	result := HelloWorld("Bagaswara")

	if result != "Hello Bagaswara" {
		// error
		t.FailNow()
	}
	// kode program ini tidak dieksekusi, karena sudah FailNow()
	// artinya langsung menghentikan unit testnya
	fmt.Println("Test HelloWorldOther Tidak Dieksekusi Selesai")
}
