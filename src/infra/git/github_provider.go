package git

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"gen-concept-api/domain/service"
)

type GitHubProvider struct{}

func NewGitHubProvider() service.GitProvider {
	return &GitHubProvider{}
}

func (p *GitHubProvider) GetFileContent(repoURL, path, token string) ([]byte, error) {
	// Simple transformation for GitHub public repos
	// Input: https://github.com/owner/repo or https://github.com/owner/repo.git
	// Output: https://raw.githubusercontent.com/owner/repo/main/path

	repoURL = strings.TrimSuffix(repoURL, "/")
	repoURL = strings.TrimSuffix(repoURL, ".git")

	if !strings.Contains(repoURL, "github.com") {
		return nil, fmt.Errorf("only github.com is supported in this basic provider")
	}

	parts := strings.Split(repoURL, "github.com/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid github url format")
	}

	repoPath := parts[1] // owner/repo

	// Construct Raw URL (Assuming 'main' branch for now as MVP)
	rawURL := fmt.Sprintf("https://raw.githubusercontent.com/%s/main/%s", repoPath, path)

	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return nil, err
	}

	if token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Try 'master' if main fails (and if likely not an auth error, though auth error on main will also return != 200)
		// For private repos, if token is wrong, we get 404 usually for security.

		rawURLMaster := fmt.Sprintf("https://raw.githubusercontent.com/%s/master/%s", repoPath, path)
		reqMaster, _ := http.NewRequest("GET", rawURLMaster, nil)
		if token != "" {
			reqMaster.Header.Set("Authorization", fmt.Sprintf("token %s", token))
		}

		respMaster, errMaster := client.Do(reqMaster)
		if errMaster == nil && respMaster.StatusCode == http.StatusOK {
			defer respMaster.Body.Close()
			return io.ReadAll(respMaster.Body)
		}

		return nil, fmt.Errorf("failed to fetch file: status %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
