package interpreter

type Store struct {
	innerMap map[string]int
}

type Variable struct {
	name string
}

const (
	MyInt = iota
	MyBool
)

type Value struct {
	Type int
	IntVal int
	BoolVal bool
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
	value Value
}

type Const struct {
	value Value
}

//Maths Operators
type Plus struct {
	left Expr
	right Expr
}

type Minus struct {
	left Expr
	right Expr
}

// Conditional Operators
type LessThan struct {
	left Expr
	right Expr
}

type MoreThan struct {
	left Expr
	right Expr
}

type EqualTo struct {
	left Expr
	right Expr
}

type Expr interface {
	Eval(Store) Value
}

type Program interface {
	Run(Store) Store
}

func (s Store) Read(variable string) int {
	return s.valRead(variable).IntVal
}

func (s Store) valRead(variable string) Value {
	return Value{MyInt,s.innerMap[variable], false}
}

func (s Store) Write(variable string, value Value) Store {
	copyOfMap := map[string]int{}
	for key, val := range s.innerMap {
		copyOfMap[key] = val
	}
	copyOfMap[variable] = value.IntVal
	return Store{copyOfMap}
}

func (c Const) Eval(s Store) Value {
	return c.value
}

func (p Plus) Eval(s Store) Value {
	return Value{MyInt,p.left.Eval(s).IntVal + p.right.Eval(s).IntVal, false}
}

func (m Minus) Eval(s Store) Value {
	return Value{MyInt,m.left.Eval(s).IntVal - m.right.Eval(s).IntVal,false}
}

func (p LessThan) Eval(s Store) Value {
	return Value{MyBool,-1, p.left.Eval(s).IntVal < p.right.Eval(s).IntVal}
}

func (mo MoreThan) Eval(s Store) Value {
	return Value{MyBool,-1, mo.left.Eval(s).IntVal > mo.right.Eval(s).IntVal}
}

func (eq EqualTo) Eval(s Store) Value {
	return Value{MyBool,-1, eq.left.Eval(s).IntVal == eq.right.Eval(s).IntVal}
}

func (v Variable) Eval(s Store) Value {
	return Value{MyInt,s.Read(v.name),false }
}

func (b Boolean) Eval(s Store) Value {
	return b.value
}

func (a Assignment) Run(s Store) Store {
	return s.Write(a.VariableName, a.RightHandSide.Eval(s))
}

func (sc SequentialComposition) Run(s Store) Store {
	firstStore := sc.LHS.Run(s)
	return sc.RHS.Run(firstStore)
}

func (c Conditional) Run(s Store) Store {
	if c.Guard.Eval(s).BoolVal {
		return c.IfCase.Run(s)
	} else {
		return c.ElseCase.Run(s)
	}

	return Store{}
}

//func (w While) Run(s Store) Store {
//	if w.condition.Eval(s) == 1 {
//		w.do.Run(s)
//	}
//	return s
//}

func evalE(e Expr, s Store) Value {
	return e.Eval(s)
}

func Run(p Program, s Store) Store {
	return p.Run(s)
}
