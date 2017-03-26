package generator

import (
	"io"
	"text/template"

	"github.com/mhamrah/gql"
)

func GenerateService(w io.Writer, od gql.ObjectDefinition) error {
	if err := genInterface(w, od); err != nil {
		return err
	}

	w.Write([]byte("\n"))

	if err := genService(w, od); err != nil {
		return err
	}

	return nil
}

func genInterface(w io.Writer, od gql.ObjectDefinition) error {
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
	return t.Execute(w, od)
}

func genService(w io.Writer, od gql.ObjectDefinition) error {
	t := template.Must(template.New("od_service").Funcs(funcMap).Parse(`
		type {{ lower .Name }}_impl struct {
			impl {{ title .Name }}
		}

		func New(impl {{ title .Name }}) gql.Service {
			return {{ lower .Name }}_impl{ impl: impl }
		}

		func (s {{ lower .Name }}_impl) Handlers() map[string]gql.HandlerFunc {
			return map[string]gql.HandlerFunc{
				{{ range .Fields -}}
					"{{ .Name }}": s.{{ title .Name }},
				{{- end }}
			}
		}

		{{ $s := lower .Name }}{{ range .Fields }}
			func (s {{ $s }}_impl) {{ title .Name }}(ctx context.Context, operation gql.Selection) (gql.Selectable, error) {
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

	return t.Execute(w, od)
}
