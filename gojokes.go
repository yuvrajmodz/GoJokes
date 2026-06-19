package gojokes

import (
	"embed"
	"fmt"
	"math/rand/v2"
	"sort"
	"sync"
)

//go:embed categories/*.json
var categoryFS embed.FS

type jokeStore struct {
	categories map[string][]string
}

var (
	defaultStore *jokeStore
	storeOnce   sync.Once
	storeErr    error
)

func getStore() (*jokeStore, error) {
	storeOnce.Do(func() {
		defaultStore, storeErr = load(categoryFS)
	})
	return defaultStore, storeErr
}

func Random() (string, error) {
	store, err := getStore()
	if err != nil {
		return "", err
	}

	all := make([]string, 0)
	for _, jokes := range store.categories {
		all = append(all, jokes...)
	}

	if len(all) == 0 {
		return "", fmt.Errorf("no jokes available")
	}

	return all[rand.IntN(len(all))], nil
}

func Category(name string) (string, error) {
	store, err := getStore()
	if err != nil {
		return "", err
	}

	jokes, ok := store.categories[name]
	if !ok {
		return "", fmt.Errorf("category not found: %q", name)
	}

	if len(jokes) == 0 {
		return "", fmt.Errorf("no jokes available in category %q", name)
	}

	return jokes[rand.IntN(len(jokes))], nil
}

func Categories() []string {
	store, err := getStore()
	if err != nil {
		return nil
	}

	names := make([]string, 0, len(store.categories))
	for name := range store.categories {
		names = append(names, name)
	}

	sort.Strings(names)
	return names
}