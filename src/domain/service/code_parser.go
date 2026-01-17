package service

type ParsedMetadata struct {
	Name   string
	Fields []ParsedField
}

type ParsedField struct {
	Name string
	Type string
}

type CodeParser interface {
	Parse(content string) (ParsedMetadata, error)
}
