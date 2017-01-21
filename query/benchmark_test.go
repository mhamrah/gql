package query

// import (
// 	"reflect"
// 	"testing"

// 	"github.com/mhamrah/gql/example/starwars"

// 	"github.com/mhamrah/gql/ast"
// )

// func BenchmarkNoReflect(b *testing.B) {
// 	foo := "foo"
// 	h := &starwars.Human{Name: &foo}
// 	result := make(map[string]interface{})

// 	for i := 0; i < b.N; i++ {
// 		result["name"] = h.Name
// 	}
// }

// func BenchmarkReflect(b *testing.B) {
// 	foo := "foo"
// 	h := &starwars.Human{Name: &foo}
// 	result := make(map[string]interface{})
// 	r := reflect.ValueOf(h).Elem()
// 	for i := 0; i < b.N; i++ {
// 		result["name"] = r.FieldByName("Name").Elem()
// 	}
// }

// func BenchmarkFieldByName(b *testing.B) {
// 	// foo := "foo"
// 	// h := &starwars.Human{Name: &foo}
// 	// result := make(map[string]interface{})

// 	// for i := 0; i < b.N; i++ {
// 	// 	//result["name"] = h.FieldByName("name")
// 	// }
// }

// var s Field

// func BenchmarkReturnStruct(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		s = wraps(i)
// 	}
// }

// var p *Field

// func BenchmarkReturnPointer(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		p = wrapp(i)
// 	}
// }

// func wraps(i int) Field {
// 	s := ReturnStruct(i)
// 	s.Ix = s.Ix - 1
// 	return s
// }
// func wrapp(i int) *Field {
// 	p := ReturnPointer(i)
// 	p.Ix = p.Ix - 1
// 	return p
// }
// func ReturnStruct(i int) Field {
// 	return Field{
// 		Name: "foo",
// 		Ix:   i,
// 	}
// }

// func ReturnPointer(i int) *Field {
// 	return &Field{
// 		Name: "foo",
// 		Ix:   i,
// 	}
// }
