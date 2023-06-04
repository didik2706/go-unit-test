package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
* Beberapa function yang digunakan untuk mengagalkan test
* 1. Fail - akan menampilkan error pada saat unit test gagal pada function test tersebut
* 2. FailNow - akan menghentikan pngetesan secara langsung
* 3. Error - sama seperti Fail tapi ada keterangan saat gagal dalam test
* 4. Fatal - sama seperti FailNow tapi ada keterangan saat gagal dalam test
*
* Cara menjalankan test
* 1. go test -v atau go test ./... | untuk menjalankan semua unit test
* 2. go test -v -run TestSubTest | untuk menjalankan spesifik function test
* 3. go test -v -run TestSubTest/Didik | untuk menjalankan spesifik sub test
*
* Untuk menjalankan benchmark
* 1. go test -v -bench=. | untuk menjalankan semua benchmark sekaligus menjalankan unit test
* 2. go test -v -run TestYgGkAda -bench=. | untuk menjalankan semua benchmark tanpa unit test
* 3. go test -v -bench=BenchmarkHelloWorld | untuk menjalankan spesifik benchmark
* 4. go test -v -bench=. ./... | untuk menjalankan semua unit test dan benchmark yang ada di project
*/

func BenchmarkHelloWorldTable(b *testing.B) {
	benchs := []struct {
		name string
		request string
	} {
		{
			name: "HelloWorld(Didik)",
			request: "Didik",
		},
		{
			name: "HelloWorld(Hidayat)",
			request: "Hidayat",
		},
	}

	for _, bench := range benchs {
		b.Run(bench.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bench.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Didik", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Didik")
		}
	})

	b.Run("Nur", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Nur")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Didik")
	}
}

func BenchmarkHelloWorldHidayat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Hidayat")
	}
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct{
		name     string
		request  string
		expected string
	} {
		{
			name: "Didik",
			request: "Didik",
			expected: "Hello, Didik",
		},
		{
			name: "Nur",
			request: "Nur",
			expected: "Hello, Nur",
		},
		{
			name: "Hidayat",
			request: "Hidayat",
			expected: "Hello, Hidayat",
		},
		{
			name: "Anies",
			request: "Anies",
			expected: "Hello, Anies",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T)() {
	t.Run("Didik", func(t *testing.T) {
		result := HelloWorld("Didik")
		assert.Equal(t, "Hello, Didik", result, "result must be 'Hello, Didik'")
	})

	t.Run("Hidayat", func(t *testing.T) {
		result := HelloWorld("Hidayat")
		assert.Equal(t, "Hello, Hidayat", result, "result must be 'Hello, Hidayat'")
	})
}

func TestMain(m *testing.M)() {
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	fmt.Println("BEFORE UNIT TEST")
}

func TestSkip(t *testing.T)() {
	if runtime.GOOS == "linux" {
		t.Skip("Cannot run in linux")
	}

	result := HelloWorld("Didik")
	assert.Equal(t, "Hello, Didik", result, "result must be 'Hello, Didik'")
}

func TestHelloWorldRequire(t *testing.T)() {
	result := HelloWorld("Didik");
	require.Equal(t, "Hello, Didik", result, "result must be 'Hello, Didik'")
}

func TestHelloWorldAssert(t *testing.T)  {
	result := HelloWorld("Didik")
	assert.Equal(t, "Hello, Didik", result, "result must be 'Hello, Didik'")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Didik");

	if result != "Hello, Didik" {
		t.Error("result must be 'Hello, Didik'")
	}

}

func TestHelloWorldSecond(t *testing.T) {
	result := HelloWorld("Didik Nur Hidayat");

	if result != "Hello, Didik Nur Hidayat" {
		t.Fatal("result must be 'Hello, Didik Nur Hidayat'")
	}

}