package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	data := make(map[string]string)

	err := filepath.Walk(".", func(path string, v os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.Contains(path, "vendor") ||
			strings.Contains(path, ".git") ||
			strings.Contains(path, "bin") ||
			strings.Contains(path, "doc") ||
			strings.Contains(path, ".idea") {
			return nil
		}
		fmt.Printf("current path: %v, file name: %v, size: %v\n", path, v.Name(), v.Size())

		content, err := ioutil.ReadFile(path)
		if err != nil {
			log.Printf("Fail to read file: %v\n", err)
		}
		fset := token.NewFileSet()

		f, err := parser.ParseFile(fset, v.Name(), string(content), 0)
		if err != nil {
			log.Printf("Fail to parse file: %v\n", err)
		}

		ast.Inspect(f, func(node ast.Node) bool {
			call, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}
			fn, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			pack, ok := fn.X.(*ast.Ident)
			if !ok {
				return true
			}
			if pack.Name != "i18n" {
				return true
			}
			if len(call.Args) == 0 {
				return true
			}
			str, ok := call.Args[0].(*ast.BasicLit)
			if !ok {
				return true
			}

			fmt.Printf("current file content: %v", str.Value)
			data[str.Value] = str.Value
			return true
		})

		return nil
	})
	if err != nil {
		log.Printf("Fail to list file: %v\n", err)
	}

	content, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("pkg/translations/en_US/data.json", content, 0664)
	if err != nil {
		log.Fatal(err)
	}
}
