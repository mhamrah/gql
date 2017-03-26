package gql

import (
	"fmt"
	"reflect"
)

func ValueAsInt(input reflect.Value) (int, error) {
	if input.Kind() != reflect.Int {
		return 0, fmt.Errorf("%v not an int", input.Kind())
	}
	return int(input.Int()), nil
}

func ValueAsString(input reflect.Value) (string, error) {
	if input.Kind() != reflect.String {
		return "", fmt.Errorf("%v not a string", input.Kind())
	}
	return input.String(), nil
}
