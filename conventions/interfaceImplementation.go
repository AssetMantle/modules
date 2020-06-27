package conventions

type abcXyz struct {
	A interface{}
	B interface{}
	C interface{}
}

var _ AbcXyz = (*abcXyz)(nil)

func (abcXyz abcXyz) GetA() interface{} { return abcXyz.A }

func (abcXyz abcXyz) GetB() interface{} { return abcXyz.B }

func (abcXyz abcXyz) GetC() interface{} { return abcXyz.C }

func (abcXyz abcXyz) DoSomething() {}

func NewAbcXyz(a interface{}, b interface{}, c interface{}) AbcXyz {
	return &abcXyz{
		A: a,
		B: b,
		C: c,
	}
}
