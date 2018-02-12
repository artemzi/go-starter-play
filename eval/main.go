package eval

import (
	"fmt"
	"math"
)

// Expr арифметическое выражение
type Expr interface {
	Eval(env Env) float64
}

// Var переменная
type Var string

type literal float64

type unary struct {
	op rune // + or -
	x  Expr
}

type binary struct {
	op   rune // +, -, * or /
	x, y Expr
}

type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}

// Env отображение переменной на значение
type Env map[Var]float64

// Eval Var возвращает nil, если переменная не определена
func (v Var) Eval(env Env) float64 {
	return env[v]
}

// Eval literal возвращает значение литерала
func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

// Eval unary вычисляет операнды и применяет к ним операцию
func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("неподдерживаемый унарный оператор: %q", u.op))
}

// Eval binary вычисляет операнды и применяет к ним операцию
func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.x.Eval(env)
	case '-':
		return b.x.Eval(env) - b.x.Eval(env)
	case '*':
		return b.x.Eval(env) * b.x.Eval(env)
	case '/':
		return b.x.Eval(env) / b.x.Eval(env)
	}
	panic(fmt.Sprintf("неподдерживаемый бинарный оператор: %q", b.op))
}

// Eval call вычисляет аргументы и вызывает соответствующую ф-ию передавая их
func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("неподдерживаемый вызов функции: %s", c.fn))
}
