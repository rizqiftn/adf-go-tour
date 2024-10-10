package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// A.58 1 - 3 Testing

var (
	kubus              Kubus   = Kubus{4}
	volumeSeharusnya   float64 = 64
	luasSeharusnya     float64 = 96
	kelilingSeharusnya float64 = 48
)

// func TestHitungVolume(t *testing.T) {
// 	t.Logf("Volume: %.2f", kubus.Volume())

// 	if kubus.Volume() != volumeSeharusnya {
// 		t.Errorf("SALAH! harusnya %.2f", volumeSeharusnya)
// 	}
// }

// func TestHitungLuas(t *testing.T) {
// 	t.Logf("Luas:%.2f", kubus.Luas())

// 	if kubus.Luas() != luasSeharusnya {
// 		t.Errorf("SALAH! harusnya %.2f", luasSeharusnya)
// 	}
// }

// func TestHitungKeliling(t *testing.T) {
// 	t.Logf("Keliling: %.2f", kubus.Keliling())

// 	if kubus.Keliling() != kelilingSeharusnya {
// 		t.Errorf("SALAH! harusnya %.2f", kelilingSeharusnya)
// 	}
// }

// A.58.4 Benchmark

// func BenchmarkHitungLuas(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		kubus.Luas()
// 	}
// }

// A.58.5 Testify
func TestHitungVolume(t *testing.T) {
	assert.Equal(t, kubus.Volume(), volumeSeharusnya, "perhitungan volume salah")
}

func TestHitungLuas(t *testing.T) {
	assert.Equal(t, kubus.Luas(), luasSeharusnya, "perhitungan luas salah")
}

func TestHitungKeliling(t *testing.T) {
	assert.Equal(t, kubus.Keliling(), kelilingSeharusnya, "perhitungan keliling salah")
}
