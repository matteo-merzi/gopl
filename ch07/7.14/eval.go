// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 198.

// Package eval provides an expression evaluator.
package eval

import (
	"fmt"
	"math"
)

//!+env

type Env map[Var]float64

//!-env

//!+Eval1

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

//!-Eval1

//!+Eval2

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	case "min":
		return math.Min(c.args[0].Eval(env), c.args[1].Eval(env))
	case "max":
		return math.Max(c.args[0].Eval(env), c.args[1].Eval(env))
	case "avrg":
		return avrg(c.args[0].Eval(env), c.args[1].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

func (v variadic) Eval(env Env) float64 {
	fargs := fArgs(v.args, env)
	switch v.fn {
	case "min":
		return min(fargs...)
	case "max":
		return max(fargs...)
	case "avrg":
		return avrg(fargs...)
	}
	panic(fmt.Sprintf("unsupported variadic function call: %s", v.fn))
}

func fArgs(vargs []Expr, env Env) []float64 {
	fargs := make([]float64, 0, len(vargs))
	for _, arg := range vargs {
		fargs = append(fargs, arg.Eval(env))
	}
	return fargs
}

func min(a ...float64) float64 {
	m := math.Min(a[0], a[1])
	for i := 2; i < len(a); i++ {
		m = math.Min(m, a[i])
	}
	return m
}

func max(a ...float64) float64 {
	m := math.Max(a[0], a[1])
	for i := 2; i < len(a); i++ {
		m = math.Max(m, a[i])
	}
	return m
}

func avrg(a ...float64) float64 {
	var sum float64
	for _, num := range a {
		sum += num
	}
	return sum / float64(len(a))
}

//!-Eval2
