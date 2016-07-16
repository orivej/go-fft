package fft

import "math"

func fft(xs []complex128, sign float64) {
	n := len(xs)

	if n == 0 || n&(n-1) != 0 {
		panic("fft: xs length must be a power of two")
	}

	bitreverse(xs)

	for stride := 1; stride < n; stride *= 2 {
		unitAngle := sign * math.Pi / float64(stride)
		sin, cos := math.Sincos(unitAngle)
		unitRot := complex(cos, sin)
		for m, w := 0, 1+0i; m < stride; m, w = m+1, w*unitRot {
			for i := m; i < n; i += 2 * stride {
				delta := w * xs[i+stride]
				xs[i], xs[i+stride] = xs[i]+delta, xs[i]-delta
			}
		}
	}
}

func bitreverse(xs []complex128) {
	msb := len(xs) >> 1
	for i, j := 1, msb; j != 0; i++ {
		if i < j {
			xs[i], xs[j] = xs[j], xs[i]
		}

		m := msb
		for ; j&m != 0; m >>= 1 {
			j ^= m
		}
		j |= m
	}
}

func FFT(xs []complex128) {
	fft(xs, -1)
}

func IFFT(xs []complex128) {
	fft(xs, 1)
	f := complex(1/float64(len(xs)), 0)
	for i := range xs {
		xs[i] *= f
	}
}
