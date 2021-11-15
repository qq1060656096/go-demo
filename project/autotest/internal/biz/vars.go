package biz

type Var map[string]interface{}

type Vars struct {
	Global Var
	Session Var
	Local Var
}

func NewVar() Var {
	v := make(Var, 10)
	return v
}

func NewVars() * Vars {
	v := Vars{
		Global: NewVar(),
		Session: NewVar(),
		Local: NewVar(),
	}
	return &v
}

type SetVars map[string]string