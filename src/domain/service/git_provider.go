package service

type GitProvider interface {
	GetFileContent(repoURL, path, token string) ([]byte, error)
}
