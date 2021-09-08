package main

import (
	"fmt"
)

type Store struct {
	innerMap map[string]int
}

type Expr interface {
	Eval(Store) int
}

type Program interface {
	Run(Store) Store
}

type Assignment struct {
	VariableName string
	RightHandSide Expr
}

type SequentialComposition struct {
	LHS Program
	RHS Program
}

type Conditional struct {
	Guard Expr
	IfCase Program
	ElseCase Program
}

type While struct {
	condition Expr
	do Program
}

type Boolean struct {
	value int
}

type Const struct {
	value int
}

type Plus struct {
	left Expr
	right Expr
}

type Minus struct {
	left Expr
	right Expr
}

type Variable struct {
	name string
}

func (s Store) Read(variable string) int {
	return s.innerMap[variable]
}

func (s Store) Write(variable string, value int) Store {
	copyOfMap := Store{innerMap: map[string]int{}}
	copyOfMap = s
	copyOfMap.innerMap[variable] = value

	return copyOfMap
}

func (c Const) Eval(s Store) int {
	return c.value
}

func (p Plus) Eval(s Store) int {
	return p.left.Eval(s) + p.right.Eval(s)
}

func (m Minus) Eval(s Store) int {
	return m.left.Eval(s) - m.right.Eval(s)
}

func (v Variable) Eval(s Store) int {
	return s.Read(v.name)
}

func (b Boolean) Eval(s Store) int {
	if b.value == 0  {
		return b.value
	} else {
		return 1
	}
}

func (a Assignment) Run(s Store) Store {
	return s.Write(a.VariableName, a.RightHandSide.Eval(s))
}

func (sc SequentialComposition) Run(s Store) Store {
	firstStore := sc.LHS.Run(s)
	return sc.RHS.Run(firstStore)
}

func (c Conditional) Run(s Store) Store {
	if c.Guard.Eval(s) == 0 {
		c.IfCase.Run(s)
	} else if c.Guard.Eval(s) == 1 {
		c.ElseCase.Run(s)
	} else {
		panic("Guard must be 0 or 1")
	}
	return s
}

func (w While) Run(s Store) Store {
	if w.condition.Eval(s) == 1 {
		w.do.Run(s)
	}
	return s
}

func evalE(e Expr, s Store) int {
	return e.Eval(s)
}

func Run(p Program, s Store) Store {
	return p.Run(s)
}

func main() {

	fmt.Printf("Run called with Assignment (x := 4) == %v\n",
		Run(Assignment{"x", Const{4}},
			Store{map[string]int{}}),
	)

	fmt.Printf("Run called with SequentialComposition (x := 4) (y := 12) == %v\n",
		Run(SequentialComposition{
				Assignment{"x", Const{4}},
				Assignment{"y", Const{12}}},
				Store{map[string]int{}}),
	)

	fmt.Printf("Run called with Conditional (0 then x := 4 else y := 12) == %v\n",
	Run(Conditional{
		Const{0},
		Assignment{"x", Const{4}},
		Assignment{"y", Const{12}}},
		Store{map[string]int{}}),
		)

//	fmt.Printf("Run called with While (x < 5) increment x == %v\n",
//		Run(While{
//			Conditional{
//				Const{1},
//
//
//			},
//		}
//,
//			map[string]int{"x": 0}),
//	)
//
//	x := 0
//	While(x < 5) {
//		x += 1
//	}
//	print x

}

