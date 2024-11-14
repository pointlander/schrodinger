// Copyright 2024 The Schrodinger Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

func Trapezoidal(x, f []complex128) complex128 {
	n := len(x)
	switch {
	case len(f) != n:
		panic("integrate: slice length mismatch")
	case n < 2:
		panic("integrate: input data too small")
	}

	integral := complex128(0.0)
	for i := 0; i < n-1; i++ {
		integral += 0.5 * (x[i+1] - x[i]) * (f[i+1] + f[i])
	}

	return integral
}

func Trapezoidal2(x, f []complex128, h float64) complex128 {
	n := len(f)
	x_mean := 0.5*f[0]*x[0] + 0.5*f[n-1]*x[n-1]
	for i := 2; i < n; i++ {
		x_mean = x_mean + f[i-1]*x[i-1]
	}
	x_mean = x_mean * complex(h, 0)
	return x_mean
}

func main() {
	x := []complex128{0, 1, 2, 3, 4}
	f := func(x complex128) complex128 {
		return x * x
	}
	y := []complex128{}
	for _, v := range x {
		y = append(y, f(v))
	}

	t := func(x complex128) complex128 {
		return x * x * x / 3
	}
	fmt.Println(t(4) - t(0))

	integral := Trapezoidal(x, y)
	fmt.Println("Integral:", integral)

	rmin, rmax := 0.0, 4.0
	ninter1 := 5.0
	h := (rmax - rmin) / ninter1
	integral2 := Trapezoidal2(x, y, h)
	fmt.Println("Integral2:", integral2)
}
