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
	value int
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

type Expr interface {
	Eval(Store) int
}

type Program interface {
	Run(Store) Store
}

func (s Store) Read(variable string) int {
	return s.innerMap[variable]
}

func (s Store) Write(variable string, value int) Store {
	copyOfMap := map[string]int{}
	for key, val := range s.innerMap {
		copyOfMap[key] = val
	}
	copyOfMap[variable] = value
	return Store{copyOfMap}
}

func (c Const) Eval(s Store) int {
	return c.valEval(s).IntVal
}

func (c Const) valEval(s Store) Value {
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
		return c.IfCase.Run(s)
	} else if c.Guard.Eval(s) == 1 {
		return c.ElseCase.Run(s)
	} else {
		panic("Guard must be 0 or 1")
	}
	return Store{}
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
