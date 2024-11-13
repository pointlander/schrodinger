// Copyright 2024 The Schrodinger Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"gonum.org/v1/gonum/integrate"
)

func main() {
	x := []float64{0, 1, 2, 3, 4}
	f := func(x float64) float64 {
		return x * x
	}
	y := []float64{}
	for _, v := range x {
		y = append(y, f(v))
	}

	integral := integrate.Trapezoidal(x, y)
	fmt.Println("Integral:", integral)
}
