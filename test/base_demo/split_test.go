package base_demo

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

func TestSpiltWithComplexSep(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

func TestTimeConsuming(t *testing.T) {
	if testing.Short() {
		t.Skip("short 模式下会跳过该测试用例")
	}
}

func TestXXX(t *testing.T) {
	t.Run("case1", func(t *testing.T) {})
	t.Run("case2", func(t *testing.T) {})
	t.Run("case3", func(t *testing.T) {})
}

func TestSplitAll(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input string
		sep   string
		want  []string
	}{
		{"base case", "a:b:c", ":", []string{"a", "b", "c"}},
		{"wrong sep", "a:b:c", ",", []string{"a:b:c"}},
		{"more sep", "abcd", "bc", []string{"a", "d"}},
		{"leading sep", "沙河有沙又有河", "沙", []string{"", "河有", "又有河"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Split(tt.input, tt.sep)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected: %#v, got: %#v", tt.want, got)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSomething(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(123, 123, "they should be equal")
	assert.NotEqual(123, 456, "they should not be equal")
	assert.Nil(nil)
	if assert.NotNil(1) {
		assert.Equal(1, 1)
	}
}
