package main

import (
	"fmt"
	gparser "go/parser"
	"go/printer"
	"go/token"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mhamrah/gql/generator"
	"github.com/mhamrah/gql/parser"
	"github.com/mhamrah/gql/validator"
)

//go:generate go run main.go sample/sample.graphql

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		log.Fatal("gqlgen requires a graphql input file")
	}

	filename := args[0]

	dir := filepath.Dir(filename)
	name := strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	err = Generate(dir, name, f)
	if err != nil {
		log.Fatal(err)
	}
}

func Generate(dir, name string, schema io.Reader) error {
	b, err := ioutil.ReadAll(schema)
	if err != nil {
		return err
	}

	doc, err := parser.ParseBytes(b)
	if err != nil {
		return err
	}

	s, err := validator.BuildSchema(doc.Schema)
	if err != nil {
		return err
	}

	files, err := generator.Generate(s, name)
	if err != nil {
		return err
	}

	genPath := fmt.Sprintf("%v/%v", dir, name)
	_ = os.MkdirAll(genPath, os.ModePerm)

	for k, v := range files {
		fset := token.NewFileSet()
		f, err := gparser.ParseFile(fset, name, v.Bytes(), gparser.ParseComments)
		if err != nil {
			log.Println("zzzzzzzzzzzz")
			log.Println(string(v.Bytes()))
			log.Fatal(err)
		}
		cfg := printer.Config{
			Mode:     printer.UseSpaces | printer.TabIndent,
			Tabwidth: 8,
		}

		out, err := os.Create(fmt.Sprintf("%v/%v", genPath, k))
		if err != nil {
			return err
		}
		if err := cfg.Fprint(out, fset, f); err != nil {
			return err
		}
	}
	return nil
}
