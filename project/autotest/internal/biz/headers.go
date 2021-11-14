package biz

type Headers map[string]string

func NewHeaders() *Headers {
	t := make(Headers, 100)
	return &t
}