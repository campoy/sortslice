package sortslice

import (
	"math/cmplx"
	"reflect"
	"sort"
)

type Int []int

func (s Int) Len() int           { return len(s) }
func (s Int) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Int) Less(i, j int) bool { return s[i] < s[j] }

type Int8 []int8

func (s Int8) Len() int           { return len(s) }
func (s Int8) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Int8) Less(i, j int) bool { return s[i] < s[j] }

type Int16 []int16

func (s Int16) Len() int           { return len(s) }
func (s Int16) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Int16) Less(i, j int) bool { return s[i] < s[j] }

type Int32 []int32

func (s Int32) Len() int           { return len(s) }
func (s Int32) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Int32) Less(i, j int) bool { return s[i] < s[j] }

type Int64 []int64

func (s Int64) Len() int           { return len(s) }
func (s Int64) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Int64) Less(i, j int) bool { return s[i] < s[j] }

type Float32 []float32

func (s Float32) Len() int           { return len(s) }
func (s Float32) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Float32) Less(i, j int) bool { return s[i] < s[j] }

type Float64 []float64

func (s Float64) Len() int           { return len(s) }
func (s Float64) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Float64) Less(i, j int) bool { return s[i] < s[j] }

type Complex64 []complex64

func (s Complex64) Len() int      { return len(s) }
func (s Complex64) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Complex64) Less(i, j int) bool {
	return cmplx.Abs(complex128(s[i])) < cmplx.Abs(complex128(s[j]))
}

type Complex128 []complex128

func (s Complex128) Len() int           { return len(s) }
func (s Complex128) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Complex128) Less(i, j int) bool { return cmplx.Abs(s[i]) < cmplx.Abs(s[j]) }

type String []string

func (s String) Len() int           { return len(s) }
func (s String) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s String) Less(i, j int) bool { return s[i] < s[j] }

type Bool []bool

func (s Bool) Len() int           { return len(s) }
func (s Bool) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Bool) Less(i, j int) bool { return !s[i] && s[j] }

func Any(s interface{}) sort.Interface {
	switch v := s.(type) {
	case []int:
		return Int(v)
	case []int8:
		return Int8(v)
	case []int16:
		return Int16(v)
	case []int32:
		return Int32(v)
	case []int64:
		return Int64(v)
	case []float32:
		return Float32(v)
	case []float64:
		return Float64(v)
	case []complex64:
		return Complex64(v)
	case []complex128:
		return Complex128(v)
	case []string:
		return String(v)
	case []bool:
		return Bool(v)
	default:
		return reflected(s)
	}
}

type reflectSlice struct{ v reflect.Value }

func (s reflectSlice) Len() int { return s.v.Len() }
func (s reflectSlice) Swap(i, j int) {
	vi, vj := s.v.Index(i), s.v.Index(j)
	aux := vi.Interface()
	vi.Set(vj)
	vj.Set(reflect.ValueOf(aux))
}

type reflectInt struct{ reflectSlice }

func (s reflectInt) Less(i, j int) bool {
	vi, vj := s.v.Index(i), s.v.Index(j)
	return vi.Int() < vj.Int()
}

type reflectFloat struct{ reflectSlice }

func (s reflectFloat) Less(i, j int) bool {
	vi, vj := s.v.Index(i), s.v.Index(j)
	return vi.Float() < vj.Float()
}

type reflectString struct{ reflectSlice }

func (s reflectString) Less(i, j int) bool {
	vi, vj := s.v.Index(i), s.v.Index(j)
	return vi.String() < vj.String()
}

type reflectBool struct{ reflectSlice }

func (s reflectBool) Less(i, j int) bool {
	vi, vj := s.v.Index(i), s.v.Index(j)
	return !vi.Bool() && vj.Bool()
}

func reflected(s interface{}) sort.Interface {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Slice {
		panic("sort only works on slices")
	}
	switch v.Type().Elem().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflectInt{reflectSlice{v}}
	case reflect.Float32, reflect.Float64:
		return reflectFloat{reflectSlice{v}}
	case reflect.String:
		return reflectString{reflectSlice{v}}
	case reflect.Bool:
		return reflectBool{reflectSlice{v}}
	}
	panic("nope")
}
