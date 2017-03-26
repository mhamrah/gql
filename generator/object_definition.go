package generator

import (
	"io"
	"text/template"

	"github.com/mhamrah/gql"
)

func GeneratObjectDefinition(w io.Writer, od gql.ObjectDefinition) error {
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
	if err := t.Execute(w, od); err != nil {
		return err
	}

	return nil
}
