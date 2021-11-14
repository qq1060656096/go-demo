package biz

type TestObject interface {
	Before()
	Execute()
	After()
	SetVars()
	Run()
	Err()
	SetErr()
}
