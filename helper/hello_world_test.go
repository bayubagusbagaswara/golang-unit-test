package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cannot trun on Windows")
	}
	result := HelloWorld("Bayu")
	require.Equal(t, "Hello Bayu", result, "Result must be 'Hello Bayu'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Bayu")
	// require akan memanggil FailNow()
	require.Equal(t, "Hello Bayu", result, "Result must be 'Hello Bayu'")
	fmt.Println("TestHelloWorld with Require Done")
}
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Bayu")
	assert.Equal(t, "Hello Bayu", result, "Result must be 'Hello Bayu'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Bayu")
	if result != "Hello Bayu" {
		t.Error("Result must be 'Hello Bayu'")
	}
	fmt.Println("Test HelloWorld Tetap Dieksekusi Selesai")
}

func TestHelloWorldOther(t *testing.T) {
	result := HelloWorld("Bagaswara")
	if result != "Hello Bagaswara" {
		t.Fatal("Result must be 'Hello Bagaswara'")
	}
	fmt.Println("Test HelloWorldOther Tidak Dieksekusi Selesai")
}
