package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Bayu")
	assert.Equal(t, "Hello Bayu", result, "Result must be 'Hello Bayu'")
	// assert ini jika error, maka akan menggagalkan testnya dan memanggil function Fail()
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Bayu")

	if result != "Hello Bayu" {
		// misal kita ganti dengan function Error dan menambahkan pesannya
		t.Error("Result must be 'Hello Bayu'")
	}
	// kode program ini tetap dieksekusi
	fmt.Println("Test HelloWorld Tetap Dieksekusi Selesai")
}

func TestHelloWorldOther(t *testing.T) {
	result := HelloWorld("Bagaswara")

	if result != "Hello Bagaswara" {
		// kita ganti dengan Fatal dan tambahkan pesan errornya
		t.Fatal("Result must be 'Hello Bagaswara'")
	}
	// kode program ini tidak dieksekusi
	fmt.Println("Test HelloWorldOther Tidak Dieksekusi Selesai")
}
