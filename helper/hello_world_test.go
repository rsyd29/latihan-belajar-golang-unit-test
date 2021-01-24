package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("World") // menjalankan function HelloWorld dan disimpan pada var result
	// pastikan return value dari func HelloWorld adalah Hello World
	// dan jangan sampai Hi World
	if result != "Hello World" { // ini akan dicek
		// apabila result tidak sama dengan Hello World
		// unit test failed and test error
		panic("Result is not 'Hello World'") // ini akan panic dan terjadi error
	}
}

func TestHelloWorldBudiman(t *testing.T) {
	result := HelloWorld("Budiman") // menjalankan function HelloWorld dan disimpan pada var result
	// pastikan return value dari func HelloWorld adalah Hello Budiman
	// dan jangan sampai Hi Budiman
	if result != "Hello Budiman" { // ini akan dicek
		// apabila result tidak sama dengan Hello Budiman
		// unit test failed and test error
		panic("Result is not 'Hello Budiman'") // ini akan panic dan terjadi error
	}
}

func TestHelloWorldRasyid(t *testing.T) {
	result := HelloWorld("Rasyid") // menjalankan function HelloWorld dan disimpan pada var result
	// pastikan return value dari func HelloWorld adalah Hello Rasyid
	// dan jangan sampai Hi Budiman
	if result != "Hello Rasyid" { // ini akan dicek
		// apabila result tidak sama dengan Hello Budiman
		// unit test failed and test error
		panic("Result is not 'Hello Rasyid'") // ini akan panic dan terjadi error
	}
}