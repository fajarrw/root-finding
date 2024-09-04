package main

import (
	"fmt"
	"math"
)

// Define the function
func f(x float64) float64 {
	return x*x*x - 6*x*x + 11*x - 6
}

// Derivative of the function
func df(x float64) float64 {
	return 3*x*x - 12*x + 11
}

// Bisection Method
func bisection(a, b, tol float64) float64 {
	if (f(a)*f(b)>0) {
		fmt.Printf("The root is not in the range [ %f, %f ]", a, b)
	}
	iter := 1
	ea := 1.0
	cold := a
	for ea > tol {
		c := (a + b) / 2
		ea = math.Abs((c-cold)/c)
		fmt.Printf("Iteration %d: a = %f, b = %f, c = %f, f(c) = %f, ea = %f\n", iter, a, b, c, f(c), ea)
		if f(c) == 0 {
			return c
		} else if f(a)*f(c) < 0 {
			b = c
		} else {
			a = c
		}
		iter++
		cold = c
	}
	return (a + b) / 2
}

// False Position Method
func falsePosition(a, b, tol float64) float64 {
	var c float64
	iter := 1
	cold := a
	ea := 1.0 
	for ea > tol {
		c = b - f(b)*(b-a)/(f(b)-f(a))
		ea = math.Abs((c-cold)/c)
		fmt.Printf("Iteration %d: a = %f, b = %f, c = %f, f(c) = %f, ea = %f\n", iter, a, b, c, f(c), ea)
		if math.Abs(f(c)) < tol {
			break
		}
		if f(c) == 0 {
			return c
		} else if f(a)*f(c) < 0 {
			b = c
		} else {
			a = c
		}
		iter++
		cold = c
	}
	return c
}

// Simple Fixed-Point Iteration
func fixedPoint(g func(float64) float64, x0, tol float64) float64 {
	x := x0
	iter := 1
	ea := 1.0
	for ea > tol {
		newX := g(x)
		ea = math.Abs((newX-x)/newX)
		fmt.Printf("Iteration %d: x = %f, g(x) = %f, ea = %f\n", iter, x, newX, ea)
		x = newX
		iter++
	}
	return x
}

// Newton-Raphson Method
func newtonRaphson(x0, tol float64) float64 {
	x := x0
	iter := 1
	ea := 1.0
	for ea > tol {
		newX := x - f(x)/df(x)
		ea = math.Abs((newX-x)/newX)
		fmt.Printf("Iteration %d: x = %f, f(x) = %f, df(x) = %f, ea = %f\n", iter, x, f(x), df(x), ea)
		x = newX
		iter++
	}
	return x
}

// Secant Method
func secant(x0, x1, tol float64) float64 {
	iter := 1
	ea := 1.0
	for ea > tol {
		newX := x1 - f(x1)*(x1-x0)/(f(x1)-f(x0))
		ea = math.Abs((newX-x1)/newX)
		fmt.Printf("Iteration %d: x0 = %f, x1 = %f, f(x1) = %f, ea = %f\n", iter, x0, x1, f(x1), ea)
		x0 = x1
		x1 = newX
		iter++
	}
	return x1
}

// Modified Secant Method
func modifiedSecant(x0, delta, tol float64) float64 {
	iter := 1
	ea := 1.0
	for ea > tol {
		fx := f(x0)
		newX := x0 - fx*delta/(f(x0+delta)-fx)
		ea = math.Abs((newX-x0)/newX)
		fmt.Printf("Iteration %d: x = %f, f(x) = %f, ea = %f\n", iter, x0, fx, ea)
		x0 = newX
		iter++
	}
	return x0
}

// Modified Newton-Raphson Method for Multiple Roots
func modifiedNewtonRaphson(x0, tol float64, m int) float64 {
	x := x0
	iter := 1
	ea := 1.0
	for ea > tol {
		
		newX := x - float64(m)*f(x)/df(x)
		ea = math.Abs((newX-x)/newX)
		fmt.Printf("Iteration %d: x = %f, f(x) = %f, df(x) = %f, ea = %f\n", iter, x, f(x), df(x), ea)
		x = newX
		iter++
	}
	return x
}

func main() {
	a, b := 0.0, 1.5 // Interval for root finding
	x0 := 0.0        // Initial guess
	tol := 0.01      // Tolerance for convergence
	delta := 0.01

	fmt.Println("Bisection Method:")
	bisection(a, b, tol)
	fmt.Println("\nFalse Position Method:")
	falsePosition(a, b, tol)
	fmt.Println("\nSimple Fixed-Point Iteration:")
	fixedPoint(func(x float64) float64 {
		return (-x*x*x + 6*x*x + 6) / 11
	}, x0, tol)
	fmt.Println("\nNewton-Raphson Method:")
	newtonRaphson(x0, tol)
	fmt.Println("\nSecant Method:")
	secant(a, b, tol)
	fmt.Println("\nModified Secant Method:")
	modifiedSecant(x0, delta, tol)
	fmt.Println("\nModified Newton-Raphson Method for Multiple Roots:")
	modifiedNewtonRaphson(x0, tol, 1)
}
