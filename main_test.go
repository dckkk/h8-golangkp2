package main_test

import (
	"fmt"
	"testing"

	main "GolangKP2"

	"github.com/stretchr/testify/assert"
)

func TestVolumeKubus(t *testing.T) {
	sisi := float64(10)
	kubus := main.Kubus{
		Sisi: sisi,
	}
	volume, _ := kubus.Volume()
	expected := float64(1000)
	assert.Equal(t, expected, volume, "testing rumus volume kubus")

	// Func ABC {A(), B(), C()} - mocking hasil A sukses, B sukses, C sukses
	// Func Check A - unit test
	// Func Validasi B - unit test
	// Func Request C - unit test
}

func TestVolumeKubusErrorSample(t *testing.T) {
	kubus := main.Kubus{
		Sisi: 4,
	}
	_, err := kubus.Volume()
	assert.NotNil(t, err)
}

func TestLuasKubus(t *testing.T) {
	sisi := float64(10)
	kubus := main.Kubus{
		Sisi: sisi,
	}
	luas := kubus.Luas()

	fmt.Println("Check 123")

	expected := float64(600)
	assert.Equal(t, expected, luas, "luas should equal with expectation")
}

func BenchmarkHitungLuas(b *testing.B) {
	kubus := main.Kubus{
		Sisi: 10,
	}
	for i := 0; i < b.N; i++ {
		kubus.Luas()
	}
}
