package service

import (
	"bytes"
	"gen-concept-api/domain/model"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
)

type HarvesterService struct{}

func NewHarvesterService() *HarvesterService {
	return &HarvesterService{}
}

func (s *HarvesterService) ScanRepo(codeContent string) ([]model.LibraryDefinition, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", codeContent, parser.ParseComments)
	if err != nil {
		// Try wrapping if it fails (e.g. snippet) - similar to Importer logic
		// But Harvester usually scans full files.
		// For MVP, if error, we might return it or try wrap.
		src := "package main\n" + codeContent
		f, err = parser.ParseFile(fset, "", src, parser.ParseComments)
		if err != nil {
			return nil, err
		}
	}

	var definitions []model.LibraryDefinition
	packageName := f.Name.Name

	ast.Inspect(f, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok {
			if fn.Name.IsExported() {
				// Reconstruct signature
				var buf bytes.Buffer
				if err := format.Node(&buf, fset, fn.Type); err == nil {
					sig := buf.String()

					desc := ""
					if fn.Doc != nil {
						desc = fn.Doc.Text()
					}

					definitions = append(definitions, model.LibraryDefinition{
						PackageName:  packageName,
						FunctionName: fn.Name.Name,
						Signature:    sig,
						Description:  desc,
						Tags:         []string{"exported"}, // Default tag
						// RepoURL would be set by caller usually
					})
				}
			}
		}
		return true
	})

	return definitions, nil
}
