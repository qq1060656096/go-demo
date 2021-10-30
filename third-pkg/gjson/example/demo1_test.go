package example

import (
	"github.com/tidwall/gjson"
	"testing"
)

func TestDemo1(t *testing.T) {
	const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	value := gjson.Get(json, "name.last")
	println(value.String())
}