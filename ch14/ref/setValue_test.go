package ref

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSetValue(t *testing.T) {
	i := 10
	iv := reflect.ValueOf(&i)
	fmt.Println("iv", iv, iv.Type())

	ivv := iv.Elem()
	fmt.Println("ivv", ivv, ivv.Type())
	ivv.SetInt(20)
	assert.Equal(t, 20, i)
}

func Test_TypeOf(t *testing.T) {
	var stringType = reflect.TypeOf((*string)(nil)).Elem()
	fmt.Println(stringType)
	assert.Equal(t, reflect.String, stringType.Kind())
	var stringSliceType = reflect.TypeOf([]string(nil))
	fmt.Println(stringSliceType)
	assert.Equal(t, reflect.String, stringSliceType.Elem().Kind())
}

func Test_NoValue(t *testing.T) {
	var str string
	assert.False(t, hasNoValue(str))
	var i interface{}
	assert.True(t, hasNoValue(i))
	n := 20
	nv := reflect.ValueOf(&n)
	assert.False(t, hasNoValue(nv))
	var m map[string]string
	assert.True(t, hasNoValue(m))
}

func hasNoValue(i interface{}) bool {
	iv := reflect.ValueOf(i)
	if !iv.IsValid() {
		return true
	}
	switch iv.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface:
		return iv.IsNil()
	default:
		return false
	}
}
