package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("World") // menjalankan function HelloWorld dan disimpan pada var result
	// pastikan return value dari func HelloWorld adalah Hello World
	// dan jangan sampai Hi World
	if result != "Hello World" { // ini akan dicek
		// apabila result tidak sama dengan Hello World
		// unit test failed and test error
		// panic("Result is not 'Hello World'") // ini akan panic dan terjadi error

	}
	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldBudiman(t *testing.T) {
	result := HelloWorld("Budiman") // menjalankan function HelloWorld dan disimpan pada var result
	// pastikan return value dari func HelloWorld adalah Hello Budiman
	// dan jangan sampai Hi Budiman
	if result != "Hello Budiman" { // ini akan dicek
		// apabila result tidak sama dengan Hello Budiman
		// unit test failed and test error
		//panic("Result is not 'Hello Budiman'") // ini akan panic dan terjadi error
		// menggunakan func Fail()
		//t.Fail()
		// menggunakan func Error()
		t.Error("Result must be 'Hello Budiman'")
	}
	// kode program ini akan tetap dieksekusi karena menggunakan func t.Fail()
	fmt.Println("TestHelloWorldBudiman Done")
}

func TestHelloWorldRasyid(t *testing.T) {
	result := HelloWorld("Rasyid") // menjalankan function HelloWorld dan disimpan pada var result
	// pastikan return value dari func HelloWorld adalah Hello Rasyid
	// dan jangan sampai Hi Budiman
	if result != "Hello Rasyid" { // ini akan dicek
		// apabila result tidak sama dengan Hello Budiman
		// unit test failed and test error
		//panic("Result is not 'Hello Rasyid'") // ini akan panic dan terjadi error
		// menggunakan func FailNow()
		//t.FailNow()
		// menggunakan func Error()
		t.Fatal("Result must be 'Hello Rasyid'")
	}
	// kode program ini tidak akan dieksekusi karena menggunakan func t.FailNow()
	fmt.Println("TestHelloWorldRasyid Done")
}

/**
Lebih disarankan menggunakan Error() ataupun Fatal() dibandingkan dengan Fail() dan
FailNow() kenapa karena kita bisa menambahkan informasi kenapa unit test kita gagal
*/

// Belajar Assertion (dengan menggunakan library testify)
/**
Assert vs Require
1. Testify menyediakan dua package untuk assertion, yaitu assert dan require
2. Saat kita menggunakan assert, jika pengecekkan gagal, maka assert akan memanggil
Fail(), artinya eksekusi unit test akan tetap dilanjutkan
3. Sedangkan jika kita menggunakan require, jika pengecekkan gagal, maka require akan
memanggil FailNow(), artinya eksekusi unit test tidak akan dilanjutkan
*/
func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Budiman")
	// Parameter Equal berisikan testing, expected, actual, msgAndArgs
	// testing itu t
	// expected hasil yang akan diharapkan
	// actual -> result atau hasil yang dijalankan
	// isi pesan apabila gagal
	assert.Equal(t, "Hello Budiman", result, "Result must be 'Hello Budiman'")
	// apabila gagal dia akan memanggil Func Fail() yang ada pada t testing
	// dia akan memanggil func Fail() maka kode program dibawah tetap bisa dijalankan
	fmt.Println("TestHelloWorld with Assertion is Done")
	// apabila gagal ini menariknya menggunakan testify
	// informasinya lebih detail
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Rasyid")
	require.Equal(t, "Hello Rasyid", result, "Result must be 'Hello Rasyid'")
	// setelah kode program diatas dijalankan maka kode dibawah tidak akan dijalankan
	// Kenapa? karena require itu memanggil FailNow()
	fmt.Println("TestHelloWorld with Require is Done")
}

// NOTE: Jadi jangan pakai if else lagi, lebih baik pakai library assertion saja

// Skip Test
func TestSkip(t *testing.T) {
	// kondisi dimana terdapat keadaan tertentu maka
	if runtime.GOOS == "linux" { // apabila os yang digunakan adalah linux
		// kita perlu skip
		t.Skip("Can't run on Linux") // maka tidak bisa berjalan di os linux
	}
	// apabila os selain linux, maka menjalankan perintah di bawah ini
	result := HelloWorld("Budiman")
	require.Equal(t, "Hello Budiman", result, "Result must be 'Hello Budiman'")
}

//TestMain
// NOTE: Ingat ini hanya berjalan pada 1 package saja, di package lain harus buat lagi TestMain
func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run() // eksekusi semua unit test

	// after
	fmt.Println("AFTER UNIT TEST")
}

//Sub Test
func TestSubTest(t *testing.T) {
	// jadi di sini ada 2 sub test
	// Sub Test Pertama
	t.Run("Budiman", func(t *testing.T) {
		// sub test ini akan melakukan kode program dengan function lain
		result := HelloWorld("Budiman")
		require.Equal(t, "Hello Budiman", result, "Result must be 'Hello Budiman'")
	})

	// Sub Test Kedua
	t.Run("Rasyid", func(t *testing.T) {
		// sub test ini akan melakukan kode program dengan function lain
		result := HelloWorld("Rasyid")
		require.Equal(t, "Hello Rasyid", result, "Result must be 'Hello Rasyid'")
	})
}

// Table Test
func TestTableHelloWorld(t *testing.T) {
	tests := []struct { // slice struct
		name     string
		request  string
		expected string
	}{
		// ini data data yang akan digunakan
		{
			name:     "HelloWorld(Budiman Rasyid)",
			request:  "Budiman Rasyid",
			expected: "Hello Budiman Rasyid",
		},
		{
			name:     "HelloWorld(Zainuddin)",
			request:  "Zainuddin",
			expected: "Hello Zainuddin",
		},
		{
			name:     "HelloWorld(Nuratitin)",
			request:  "Nuratitin",
			expected: "Hello Nuratitin",
		},
		{
			name:     "HelloWorld(Subhan Amin)",
			request:  "Subhan Amin",
			expected: "Hello Subhan Amin",
		},
		// di sini kalau ingin menambahkan datanya sama seperti data di atas
	}

	// Melakukan Iterasi
	for _, test := range tests {
		// memanggil function sub test
		t.Run(test.name, func(t *testing.T) {
			// parameter didapatkan dari slice
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// Kode: Membuat Benchmark Function
/**
Menjalankan Benchmark
1. Untuk menjalankan seluruh benchmark di module/package, kita bisa menggunakan
perintah sama seperti test, namun ditambahkan parameter bench:
go test -v -bench=.
2. Jika kita hanya ingin menjalankan benchmark tanpa unit test, kita bisa gunakan
perintah:
go test -v -run=NotMatchUnitTest -bench=.
3. Kode diatas selain menjalankan benchmark, akan menjalankan unit test juga, jika
kita hanya ingin menjalankan benchmark tertentu, kita bisa gunakan perintah:
go test -v -run=NotMatchUnitTest -bench=BenchmarkTest
4. Jika kita menjalankan benchmark di root module/package dan ingin semua module/
package dijalankan, kita bisa gunakan perintah:
go test -v -bench=. ./...
*/
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("World")
	}
}

func BenchmarkHelloWorldRasyid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Rasyid")
	}
}

// Sub Benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("World", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("World")
		}
	})
	b.Run("Rasyid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Rasyid")
		}
	})
}

// Table Benchmark
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Budiman",
			request: "Budiman",
		},
		{
			name:    "Rasyid",
			request: "Rasyid",
		},
		{
			name:    "BudimanRasyidZainuddin",
			request: "Budiman Rasyid Zainuddin",
		},
		{
			name:    "SubhanAminZainuddin",
			request: "Subhan Amin Zainuddin",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			HelloWorld(benchmark.request)
		})
	}
}
