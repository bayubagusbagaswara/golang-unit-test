# Belajar Golang Unit Test

## Agenda

- Pengenalan Software Testing
- Testing Package
- Unit Test
- Assertion
- Mock, dan
- Benchmark

## Test Pyramid

- UI Testing atau End-to-End testing
- Service Testing atau Integration Testing
- Unit Testing

## Unit Test

- Unit test akan fokus menguji bagian kode program terkecil, biasanya menguji sebuah method
- Unit test biasanya dibuat kecil dan cepat. Oleh karena itu, biasanya unit test lebih banyak dari kode program aslinya, karena semua skenario pengujian akan dicoba di unit test
- Unit test bisa digunakan sebagai cara untuk meningkatkan kualitas kode program kita

## Pengenalan Testing Package

- Di bahasa pemrograman lain, biasanya untuk implementasi unit test, kita butuh library atau framework khusus untuk unit test. Misal di Java, butuh JUnit
- Berbeda dengan Go-Lang, di Go-Lang untuk unit test sudah disediakan sebuah package khusus bernama `testing`
- Selain itu, untuk menjalankan unit test, di Go-Lang juga sudah disediakan perintahnya
- Hal ini membuat implementasi unit testing di golang sangat mudah dibanding dengan bahasa pemrograman yang lain
- https://golang.org/pkg/testing/

## testing.T

- Go-Lang menyediakan sebuah `struct` yang bernama `testing.T`
- Struct ini digunakan `untuk unit test` di Go-Lang

## testing.M

- testing.M adalah struct yang disediakan Go-Lang `untuk mengatur lif cycle testing`

## testing.B

- testing.B adalah `struct` yang disediakan Go-Lang `untuk melakukan benchmarking`
- Di Go-Lang untuk melakukan benchmark (mengukur kecepatan kode program) pun sudah disediakan, sehingga kita tidak perlu menggunakan library atau framework tambahan

# Membuat Unit Test

## Aturan File Test

- Go-Lang memiliki aturan cara membuat file khusus untuk unit test
- Untuk membuat file unit test, kita harus menggunakan akhiran `_test`
- Jadi misalnya kita membuat file bernama hello_world.go, artinya untuk membuat unit testnya, kita harus membuat file hello_world_test.go

## Aturan Function Unit Test

- Selain aturan nama file, fi Go-Lang juga sudah diatur untuk nama function unit test
- `Nama function` untuk unit test harus `diawali dengan nama Test`
- Misal jika kita ingin mengetest function HelloWorld, maka kita akan membuat function unit test dengan nama TestHelloWorld
- Tak ada aturan untuk nama belakang function unit test harus sama dengan nama function yang akan di test, yang penting adalah harus diawalin dengan kata Test
- Selanjutnya harus memiliki parameter `(t *testing.T)` dan tidak mengembalikan return value

## Menjalankan Unit Test

- Untuk menjalankan unit test kita bisa menggunakan perintah : `go test`
- Jika kita ingin lebih detail function test apa saja yang sudah di running, kita bisa gunakan perintah : `go test -v`
- Dan jika kita ingin memilih function unit test mana yang ingin di running, kita bisa gunakan perintah : `go test -v -run TestNamaFunction`

## Menjalankan Semua Unit Test

- Jika kita ingin menjalankan semua unit test dari top folder module nya, kita bisa gunakan perintah : `go test./...`

## Menggagalkan Unit Test

- Menggagalkan unit test menggunakan panic bukanlah hal yang bagus
- Go-Lang sendiri sudah menyediakan cara untuk mengagalkan unit test menggunakan `testing.T`
- Terdapat function Fail(), FailNow(), Error() dan Fatal() jika kita ingin menggagalkan unit test

### t.Fail() dan t.FailNow()

- Terdapat dua function untuk menggagalkan unit test, yaitu Fail() dan FailNow()
- Fail() akan menggagalkan unit test. Namun tetap melanjutkan eksekusi unit testnya. Akan tetapi, ketika diakhir (saat test selesai), maka unit test tersebut dianggap gagal
- FailNow() akan menggagalkan unit test saat ini juga, tanpa melanjutkan eksekusi unit test
