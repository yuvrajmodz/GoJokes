package gojokes

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

func load(fsys fs.FS) (*jokeStore, error) {
	store := &jokeStore{
		categories: make(map[string][]string),
	}

	entries, err := fs.ReadDir(fsys, "categories")
	if err != nil {
		return nil, fmt.Errorf("failed to load categories: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if filepath.Ext(name) != ".json" {
			continue
		}

		category := strings.TrimSuffix(name, ".json")

		data, err := fs.ReadFile(fsys, "categories/"+name)
		if err != nil {
			return nil, fmt.Errorf("failed to read category %q: %w", category, err)
		}

		var jokes []string
		if err := json.Unmarshal(data, &jokes); err != nil {
			return nil, fmt.Errorf("failed to parse category %q: %w", category, err)
		}

		if len(jokes) > 0 {
			store.categories[category] = jokes
		}
	}

	if len(store.categories) == 0 {
		return nil, fmt.Errorf("failed to load categories: no joke categories found")
	}

	return store, nil
}