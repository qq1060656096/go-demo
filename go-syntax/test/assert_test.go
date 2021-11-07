package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestAssertOk(t *testing.T) {
	assert.Equal(t, "test", "test", "TestAssertOk fail")
}

func TestAssertFail(t *testing.T) {
	assert.Equal(t, "test", "test fail", "TestAssertFail fail")
}

func BenchmarkStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringJoin("hello world", 100000)
	}
}

func stringJoin(str string, times int) string {
	s := ""
	for i := 0; i < times; i ++ {
		s += str + fmt.Sprintf("%d", times)
	}
	return s
}

func BenchmarkStringJoinStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringJoinStringBuilder("hello world", 100000)
	}
}

func stringJoinStringBuilder(str string, times int) string {
	s := strings.Builder{}
	for i := 0; i < times; i ++ {
		s.WriteString(str)
		s.WriteByte(byte (times))
	}
	return s.String()
}