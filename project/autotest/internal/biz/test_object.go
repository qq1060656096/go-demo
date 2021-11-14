package biz

type TestObject interface {
	Before()
	Execute()
	After()
	ParseVars()
	Run()
	Err()
	SetErr(err error)
	Errs()
}
