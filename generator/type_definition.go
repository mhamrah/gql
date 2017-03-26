package generator

import (
	"io"

	"github.com/mhamrah/gql"
)

func GenerateTypeDefinition(w io.Writer, td gql.TypeDefinition) {
	switch t := td.(type) {
	case gql.ObjectDefinition:
		GeneratObjectDefinition(w, t)
	}
}
