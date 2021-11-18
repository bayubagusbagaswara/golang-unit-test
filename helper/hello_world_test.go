package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// kita bikin benchmark table
func BenchmarkTable(b *testing.B) {
	// bikin dulu variable benchmark untuk struct slice
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Bayu",
			request: "Bayu",
		},
		{
			name:    "Bagaswara",
			request: "Bagaswara",
		},
		{
			name:    "BayuBagusBagaswara",
			request: "BayuBagusBagaswara",
		},
	}
	// iterasi menggunakan for
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

// kita bikin Sub Benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("Bayu", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Bayu")
		}
	})

	b.Run("Bagaswara", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Bagaswara")
		}
	})
}

// kita bikin function benchmark
func BenchmarkHelloWorld(b *testing.B) {

	// lakukan perulangan sebanyak N
	for i := 0; i < b.N; i++ {

		// eksekusi kode program kita disini
		// misalnya kita punya function dengan parameter Bayu
		HelloWorld("Bayu")
	}
}
func BenchmarkHelloWorldBagaswara(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Bagaswara")
	}
}

func TestTableHelloWorld(t *testing.T) {
	// kita bikin test nya
	// berupa slice of struct
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		// seperti kita mendefiniskan datanya dan ekspektasinya
		// lalu tinggal running testnya menggunakan perulangan
		{
			// masukkan nilai tiap slice nya
			// misal ini adalah function dengan parameter Bayu
			name:     "Bayu",
			request:  "Bayu",
			expected: "Hello Bayu",
		},
		{
			// misal ini adalah function dengan parameter Bagaswara
			name:     "Bagaswara",
			request:  "Bagaswara",
			expected: "Hello Bagaswara",
		},
	}
	// lakukan perulangan untuk mejalankan testnya
	for _, test := range tests {
		// jalankan testnya dan ambil index namenya
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			// hasilnya harus sesuai dengan ekspektasi kita
			require.Equal(t, test.expected, result)
		})
	}

}

func TestSubTest(t *testing.T) {
	t.Run("Bayu", func(t *testing.T) {
		// bikin sub test
		result := HelloWorld("Bayu")
		require.Equal(t, "Hello Bayu", result, "Result must be 'Hello Bayu'")
	})

	t.Run("Bagus", func(t *testing.T) {
		// bikin subtest
		result := HelloWorld("Bagus")
		require.Equal(t, "Hello Bagus", result, "Result must be 'Hello Bagus'")

	})
}

func TestMain(m *testing.M) {
	// before, sebelum unit testnya dijalankan
	fmt.Println("BEFORE UNIT TEST")

	m.Run() // eksekusi semua unit test

	// after, setelah unit testnya selesai dijalankan
	fmt.Println("AFTER UNIT TEST")
}

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
