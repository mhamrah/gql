package gql

import "fmt"

type Field struct {
	Name         string
	Arguments    map[string]Argument
	SelectionSet []Selection
	Directives   []Directive
	Alias        string
	Ix           int
}

func (f Field) StringValue(arg, defaultValue string, required bool) (string, error) {
	input, ok := f.Arguments[arg]
	if !ok && required {
		return "", fmt.Errorf("required argument %v is not present in list to %v", arg, f.Name)
	}

	if !ok && !required {
		return defaultValue, nil
	}

	if !input.Value.IsValid() {
		return defaultValue, fmt.Errorf("%v does not contain a valid value", arg)
	}

	result, err := ValueAsString(input.Value)
	if err != nil {
		return "", err
	}

	return result, nil
}
