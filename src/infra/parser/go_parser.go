package parser

import (
	"fmt"
	"gen-concept-api/domain/service"
	"go/ast"
	"go/parser"
	"go/token"
)

type GoParser struct{}

func NewGoParser() service.CodeParser {
	return &GoParser{}
}

func (p *GoParser) Parse(content string) (service.ParsedMetadata, error) {
	fset := token.NewFileSet()
	// To parse a code snippet that might be just a struct, we might need to wrap it in "package main" if it's not present.
	// But usually AST works better on full files.
	// Let's try parsing as a file, and if it fails, prepend package main.

	src := content
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		// Try wrapping
		src = "package main\n" + content
		f, err = parser.ParseFile(fset, "", src, 0)
		if err != nil {
			return service.ParsedMetadata{}, fmt.Errorf("failed to parse go code: %v", err)
		}
	}

	var metadata service.ParsedMetadata

	// Inspect AST to find the first Struct Type
	ast.Inspect(f, func(n ast.Node) bool {
		if t, ok := n.(*ast.TypeSpec); ok {
			if structType, ok := t.Type.(*ast.StructType); ok {
				metadata.Name = t.Name.Name

				for _, field := range structType.Fields.List {
					if len(field.Names) > 0 {
						for _, name := range field.Names {
							fieldName := name.Name
							fieldType := "unknown"
							if ident, ok := field.Type.(*ast.Ident); ok {
								fieldType = ident.Name
							}
							// Handle other types like selector expr (time.Time) later if needed for MVP

							metadata.Fields = append(metadata.Fields, service.ParsedField{
								Name: fieldName,
								Type: fieldType,
							})
						}
					}
				}
				return false // Stop after finding first struct (MVP assumption: one main struct per snippet)
			}
		}
		return true
	})

	if metadata.Name == "" {
		return service.ParsedMetadata{}, fmt.Errorf("no struct definition found in code")
	}

	return metadata, nil
}
