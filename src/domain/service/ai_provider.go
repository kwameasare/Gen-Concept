package service

type AIProvider interface {
	GenerateContent(prompt string) (string, error)
}
