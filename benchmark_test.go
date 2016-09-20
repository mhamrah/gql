package gql

import (
	"reflect"
	"testing"

	"github.com/mhamrah/gql/example/starwars"
)

func BenchmarkNoReflect(b *testing.B) {
	foo := "foo"
	h := &starwars.Human{Name: &foo}
	result := make(map[string]interface{})

	for i := 0; i < b.N; i++ {
		result["name"] = h.Name
	}
}

func BenchmarkReflect(b *testing.B) {
	foo := "foo"
	h := &starwars.Human{Name: &foo}
	result := make(map[string]interface{})
	r := reflect.ValueOf(h).Elem()
	for i := 0; i < b.N; i++ {
		result["name"] = r.FieldByName("Name").Elem()
	}
}

func BenchmarkFieldByName(b *testing.B) {
	foo := "foo"
	h := &starwars.Human{Name: &foo}
	result := make(map[string]interface{})

	for i := 0; i < b.N; i++ {
		result["name"] = h.FieldByName("name")
	}
}
