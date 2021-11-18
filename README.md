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

### t.Error(args...) dan t.Fatal(args...)

- Error() functin lebih seperti melakukan log(print) error. Karena dengan Error ini kita bisa menambahkan argument. Misalnya seperti pesan errornya apa ketika terjadi error
- Namun, setelah melakukan `log error`, lalu dia akan secara otomatis memanggil function Fail(), sehingga mengakibatkan unit test dianggap gagal
- Namun, karena hanya memanggil Fail(), artinya eksekusi unit test (satu unit test saja) akan tetap berjalan sampai selesai
- Fatal() mirip dengan Error(). Hanya saja, setelah melakukan log error, lalu dia akan memanggil FailNow(), sehingga mengakibatkan eksekusi unit test berhenti

## Assertion

- Sangat disarankan menggunakan Assertion untuk melakukan pengecekan
- Sayangnya, di Golang tidak menyediakan package untuk Assertion, sehingga kita butuh menambahkan library untuk melakukan Assertion ini

## Testify

- Salah satu library Assertion yang paling populer di Golang adalah Testify
- https://github.com/stretchr/testify
- tambahkan di Go Module kita : `go get github.com/stretchr/testify`

### assert vs require

- Testify menyediakan dua package untuk assertion, yaitu assert dan require
- Saat kita menggunakan `assert`. Jika pengecekannya gagal, maka assert akan memanggil function `Fail()`, dimana artinya adalah unit test akan tetap dieksekusi sampai kode program selesai
- Sedangkan jika kita menggunakan `require`. Jika pengecekan gagal, maka require akan memanggil function `FailNow()`, artinya eksekusi unit test tidak akan dilanjutkan

## Skip Test atau Membatalkan Test

- Unit test nya bakal di ignore
- Kadang dalam keadaan tertentu, kita ingin membatalkan eksekusi unit test
- Di Go-Lang juga kita bisa membatalkan eksekusi unit test
- Yakni dengan menggunakan function `Skip()`

## Before dan After Test

- Kadang kita ingin melakukan sesuatu sebelum dan setelah sebuah unit test dieksekusi
- Jikalau kode yang kita lakukan sebelum dan setelah selalu sama antar unit test function, maka membuat manual di unit test function tidaklah disarankan
- Di Golang terdapat fitur yang bernama `testing.M`
- Fitur ini bernama `Main`, dimana digunakan untuk mengatur eksekusi unit test. Namun hal ini juga bisa kita gunakan untuk melakukan Before dan After di unit test

## testing.M

- Untuk mengatur eksekusi unit test, kita cukup membuat sebuah function bernama `TestMain` dengan `parameter testing.M`
- Jika terdapat function TestMain tersebut, maka secara otomatis Golang akan mengeksekusi function ini tiap kali menjalankan unit test di sebuah package
- Dengan ini kita bisa mengatur Before dan After unit test sesuai dengan yang kita mau
- Ingat, function TestMain itu dieksekusi hanya sekali per Golang package, bukan per tiap function unit test

## Sub Test

- Golang mendukung fitur pembuatan function unit test di dalam function unit test
- Fitur ini memang sedikit aneh dan jarang sekali dimiliki di unit test di bahasa pemrograman yang lainnya
- Untuk membuat sub test, kita bisa menggunakan function `Run()`

### Menjalankan Hanya Sub Test

- Jika ingin menjalankan sebuah unit test function, kita bisa menggunakan perintah : `go test -run TestNamaFunction`
- Jika kita ingin menjalankan hanya salah satu sub test, kita bisa menggunakan perintah : `go test -run TestNamaFunction/NamaSubTest`
- Atau untuk semua function di unit test tapi hanya sub test tertentu, kita bisa gunakan perintah : `go test -run /NamaSubTest`

## Table Test

- Jika diperhatikan, sebenarnya dengan sub test, kita bisa membuat `test secara dinamis`
- Dan fitur sub test ini, biasa digunakan oleh programmer Go-Lang untuk membuat test dengan konsep `Table Test`
- Table Test yaitu dimana kita menyediakan `data berupa slice` yang berisi `parameter dan ekspektasi` hasil dari unit test
- Lalu slice tersebut kita iterasi menggunakan sub test

## Mock

- Mock adalah object yang sudah kita program dengan ekspektasi tertentu sehingga ketika dipanggil, maka dia akan menghasilkan data yang sudah kita program diawal
- Mock adalah salah satu teknik dalam unit testing. Dimana kita bisa membuat mock object dari suatu object yang memang sulit untuk di testing
- Misal kita ingin membuat unit test. Namun ternyata ada kode program kita yang harus memanggil API Call ke third party service. Hal ini sangat sulit untuk di test, karena unit testing kita harus selalu memanggil third party service. Dan belum tentu response nya sesuai dengan apa yang kita mau
- Pada kasus seperti itu, cocok sekali untuk menggunakan mock object

## Testify Mock

- Untuk membuat mock object, tidak ada fitur bawaan Go-Lang. Namun, kita bisa menggunakan library testify yang sebelumnya kita gunakan untuk assertion
- Testify mendukung pembuatan mock object, sehingga cocok untuk kita gunakan ketika ingin membuat mock object
- Namun, perlu diperhatikan, jika desain kode program kita jelek, akan sulit untuk melakukan mocking. Jadi pastikan kita melakukan pembuatan desain kode program kita dengan baik

## Contoh Kasus: Aplikasi Query ke Database

- Dimana kita akan buat layer Service sebagai business logic, dan layer Repository sebagai jembatan ke database
- Agar kode kita mudah untuk di test, disarankan agar membuat kontrak berupa Interface

## Benchmark

- Selain unit test, Go-Lang testing package juga mendukung melakukan benchmark
- Benchmark adalah mekanisme `menghitung kecepatan performa kode aplikasi kita`
- Benchmark di Go-Lang dilakukan dengan cara secara otomatis melakukan iterasi kode yang kita panggil berkali-kali sampai waktu tertentu
- Kita tidak perlu menentukan jumlah iterasi dan lamanya, karena itu sudah diatur oleh `testing.B` bawaan dari testing package

## testing.B

- testing.B adalah struct yang digunakan untuk melakukan benchmark
- testing.B mirip dengan testing.T, terdapat function Fail(), FailNow(), Error(), Fatal() dan lain-lain
- Yang membedakan, ada beberapa attribute dan function tambahan yang digunakan untuk melakukan benchmark
- Salah satunya adalah attribute N, ini digunakan untuk melakukan total iterasi sebuah benchmark

## Cara Kerja Benchmark

- Cara kerja benchmark di Go-Lang sangat sederhana
- Dimana kita hanya perlu membuat `perulangan (iterasi for)` sejumlah N attribute
- Nanti secara otomatis Go-Lang akan melakukan eksekusi sejumlah perulangan yang ditentukan secara otomatis
- Lalu mendeteksi berapa lama proses tersebut berjalan, dan disimpukan performa benchmark nya dalam waktu

## Benchmark Function

- Mirip seperti unit test, untuk benchmark pun, di Go-Lang sudah ditentukan nama function nya. Harus diawali dengan kata Benchmark, misal BenchmarkHelloWorld, BenchmarkXxx
- Selain itu, harus memiliki parameter `(b *testing.B)`
- Dan tidak boleh mengembalikan return value
- Untuk nama file benchmark, sama seperti unit test, diakhiri dengan `_test`, misal hello_world_test.go

## Menjalankan Benchmark

- Untuk menjalankan seluruh benchmark di module, kita bisa menggunakan perintah sama seperti test. Namun, ditambahkan parameter bench : `go test -v -bench=.`
- Jika kita hanya ingin menjalankan benchmark tanpa unit test, kita bisa gunakan perintah : `go test -v -run=NotMathUnitTest -bench=.`
- Kode diatas selain menjalankan benchmark, juga akan menjalanka unit test nya
- Jika kita hanya ingin menjalankan benchmark tertentu, kita bisa gunakan perintah : `go test -v -run=NotMathUnitTest -bench=BenchmarkTest`
- Jika kita menjalankan benchmark di root module dan ingin semua module dijalankan, kita bisa gunakan perintah : `go test -v -bench=. ../...`

## Sub Benchmark

- Sama seperti testing.T, di testing.B juga kita bisa membuat sub benchmark menggunakan function `Run()`

## Menjalankan Hanya Sub Benchmark

- Saat kita menjalankan benchmark function, maka semua sub benchmark akan berjalan
- Namun, jika kita ingin menjalankan salah satu sub benchmark saja, kita bisa gunakan perintah : `go test -v -bench=BenchmarkNama/NamaSub`

## Table Benchmark

- Sama seperti di unit test, programmer Go-Lang terbiasa membuat table benchmark juga
- Ini digunakan agar kita bisa mudah melakukan performance test dengan kombinasi data berbeda-beda tanpa harus membuat banyak benchmark function
