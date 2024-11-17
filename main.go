// Copyright 2024 The Schrodinger Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func Trapezoidal(x, f []float64) float64 {
	n := len(x)
	switch {
	case len(f) != n:
		panic("integrate: slice length mismatch")
	case n < 2:
		panic("integrate: input data too small")
	}

	integral := 0.0
	for i := 0; i < n-1; i++ {
		integral += 0.5 * (x[i+1] - x[i]) * (f[i+1] + f[i])
	}

	return integral
}

func Trapezoidal2(x, f []float64, h float64) float64 {
	n := len(f)
	x_mean := 0.5*f[0] + 0.5*f[n-1]
	for i := 2; i < n; i++ {
		x_mean = x_mean + f[i-1]
	}
	x_mean = x_mean * h
	return x_mean
}

func main() {
	x := []float64{0, .1, .2, .3, .4, .5, .6, .7, .8, .9}
	f := func(x float64) complex128 {
		return cmplx.Exp(complex(x, 0) * 1i)
	}
	y := []float64{}
	for _, v := range x {
		yy := f(v)
		y = append(y, real(cmplx.Conj(yy)*complex(math.Exp(-v*v), 0)*yy))
	}

	t := func(x float64) float64 {
		return .5 * math.Sqrt(math.Pi/2.0) * math.Erf(math.Sqrt(2)*x)
	}
	fmt.Println(t(.9) - t(0))

	integral := Trapezoidal(x, y)
	fmt.Println("Integral:", integral)

	rmin, rmax := 0.0, .9
	ninter1 := float64(len(x))
	h := (rmax - rmin) / ninter1
	integral2 := Trapezoidal2(x, y, h)
	fmt.Println("Integral2:", integral2)
}
