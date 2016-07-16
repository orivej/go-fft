package fft

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

const delta = 1e-15

func copyComplex128Slice(xs []complex128) []complex128 {
	ys := make([]complex128, len(xs))
	copy(ys, xs)
	return ys
}

func assertInDeltaSlice(t *testing.T, expected, actual []complex128, delta float64) {
	for i := range expected {
		assert.InDelta(t, real(expected[i]), real(actual[i]), delta)
		assert.InDelta(t, imag(expected[i]), imag(actual[i]), delta)
	}
}

func TestFFT(t *testing.T) {
	for _, test := range [][2][]complex128{
		{{2}, {2}},
		{{1, 2, 3, 4}, {10, -2 + 2i, -2, -2 - 2i}},
	} {
		xs, ft := test[0], test[1]
		ys := copyComplex128Slice(xs)
		FFT(ys)
		assertInDeltaSlice(t, ft, ys, delta)
		IFFT(ys)
		assertInDeltaSlice(t, xs, ys, delta)
	}
}

func TestPanic(t *testing.T) {
	assert.Panics(t, func() { FFT([]complex128{}) })
	assert.Panics(t, func() { FFT([]complex128{1, 2, 3}) })
}

func BenchmarkFFT(b *testing.B) {
	xs := make([]complex128, 1024)
	for i := range xs {
		f := float64(i)
		scale := math.Log(f + 1)
		xs[i] = complex(scale*math.Sin(f), scale*math.Cos(f))
	}
	for i := 0; i < b.N; i++ {
		FFT(xs)
		IFFT(xs)
	}
}
