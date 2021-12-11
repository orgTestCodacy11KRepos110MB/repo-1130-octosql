package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/Masterminds/semver"
)

var officialPluginRepositoryURL = func() string {
	if url, ok := os.LookupEnv("OCTOSQL_PLUGIN_REPOSITORY_OFFICIAL_URL"); ok {
		return url
	}
	return "https://raw.githubusercontent.com/cube2222/octosql/master/plugin_repository.json"
}()

var additionalPluginRepositoryURLs = func() []string {
	if urls, ok := os.LookupEnv("OCTOSQL_PLUGIN_REPOSITORY_ADDITIONAL_URLS"); ok {
		return strings.Split(urls, ",")
	}
	return []string{}
}()

func GetRepositories(ctx context.Context) ([]Repository, error) {
	officialRepository, err := getRepository(ctx, officialPluginRepositoryURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't get official repository: %w", err)
	}
	additionalRepositories := make([]Repository, len(additionalPluginRepositoryURLs))
	for i := range additionalPluginRepositoryURLs {
		repository, err := getRepository(ctx, additionalPluginRepositoryURLs[i])
		if err != nil {
			return nil, fmt.Errorf("couldn't get additional repository from '%s': %w", additionalPluginRepositoryURLs[i], err)
		}
		if repository.Prefix == "" {
			return nil, fmt.Errorf("repository from '%s' doesn't have a prefix, additional repositories must have prefixes", additionalPluginRepositoryURLs[i])
		}
		additionalRepositories[i] = repository
	}

	return append([]Repository{officialRepository}, additionalRepositories...), nil
}

func getRepository(ctx context.Context, url string) (Repository, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Repository{}, fmt.Errorf("couldn't create request to get plugin repository contents: %w", err)
	}
	req = req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Repository{}, fmt.Errorf("couldn't get plugin repository contents: %w", err)
	}
	defer res.Body.Close()

	var out Repository
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return Repository{}, fmt.Errorf("couldn't decode plugin repository contents: %w", err)
	}

	return out, nil
}

type Repository struct {
	Name        string
	Description string
	Prefix      string
	Plugins     []Plugin
}

type Plugin struct {
	Name string `json:"name"`

	// Description is a *short* description of the plugin.
	Description  string `json:"description"`
	ManifestURL  string `json:"manifest_url"`
	Website      string `json:"website"`
	ContactEmail string `json:"contact_email"`
	License      string `json:"license"`
}

type Manifest struct {
	BinaryDownloadURLPattern string    `json:"binary_download_url_pattern"`
	ReadmeURL                string    `json:"readme_url"`
	Versions                 []Version `json:"versions"`
}

func (m *Manifest) GetBinaryDownloadURL(version semver.Version) string {
	return strings.NewReplacer(
		"{{os}}", runtime.GOOS,
		"{{arch}}", runtime.GOARCH,
		"{{version}}", version.String(),
	).Replace(m.BinaryDownloadURLPattern)
}

type Version struct {
	Number semver.Version `json:"number"`
}