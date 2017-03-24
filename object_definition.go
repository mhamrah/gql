package gql

import (
	"io"

	"text/template"
)

type ObjectDefinition struct {
	Name        string
	Implements  []string
	Fields      []FieldDefinition
	Directives  []Directive
	IsOperation bool
}

func (o ObjectDefinition) Generate(w io.Writer) error {
	w.Write([]byte("\n"))
	t := template.Must(template.New("od_gen").Funcs(funcMap).Parse(`
		type {{ title .Name }} struct {
			{{ range .Fields -}}
				{{ title .Name }} {{ typeMap .Type.Name }}
			{{ end }}
		}

		func (x {{ title .Name }}) ValueFromName(name string) interface{} {
			switch name {
				{{ range .Fields -}}
					case "{{ .Name }}":
						return x.{{ title .Name }}
				{{ end }}
			}
			return nil
		}
	`))
	if err := t.Execute(w, o); err != nil {
		return err
	}

	return nil
}

func (o ObjectDefinition) GenerateOperation(w io.Writer) error {
	if err := o.genInterface(w); err != nil {
		return err
	}

	w.Write([]byte("\n"))

	if err := o.genService(w); err != nil {
		return err
	}

	return nil
}

func (o ObjectDefinition) genInterface(w io.Writer) error {
	t := template.Must(template.New("od_interface").Funcs(funcMap).Parse(`
		type {{ title .Name }} interface {
			{{ range .Fields -}}
				{{ title .Name }}(ctx context.Context
					{{- range .Arguments -}}
						, {{ .Name }} {{ typeMap .Type.Name -}}
					{{ end }}) ({{ typeMap .Type.Name }}, error)
			{{ end }}
		}
	`))
	return t.Execute(w, o)
}

func (o ObjectDefinition) genService(w io.Writer) error {
	t := template.Must(template.New("od_service").Funcs(funcMap).Parse(`
		type {{ lower .Name }}_impl struct {
			impl {{ title .Name }}
		}

		func New(impl {{ title .Name }}) gql.Service {
			return {{ lower .Name }}_impl{ impl: impl }
		}

		func (s {{ lower .Name }}_impl) Handlers() map[string]gql.GqlFunc {
			return map[string]gql.GqlFunc{
				{{ range .Fields -}}
					"{{ .Name }}": s.{{ title .Name }},
				{{- end }}
			}
		}

		{{ $s := lower .Name }}{{ range .Fields }}
			func (s {{ $s }}_impl) {{ title .Name }}(ctx context.Context, operation ast.Selection) (gql.NamedLookup, error) {
				{{ range $key, $arg := .Arguments }}
					{{ $init := typeInitializer $arg }}
					{{ $init.Init }}
					if input, ok := operation.Field.Arguments["{{ $key }}"]; ok {
						var err error
						if !input.Value.IsValid() {
							return nil, fmt.Errorf("%v does not contain a valid value", {{ $key }})
						}
						{{ $init.Setter }}
						if err != nil {
							return nil, err
						}
					}

				{{ end }}

				success, err := s.impl.{{ title .Name }}(ctx
					{{- range $key, $arg := .Arguments -}}
						, {{ $key }}
					{{- end -}}
				)
				if err != nil {
					return nil, err
				}
				return success, nil
			}
		{{ end }}
	`))

	return t.Execute(w, o)
}

func (o ObjectDefinition) NamedType() string {
	return o.Name
}
