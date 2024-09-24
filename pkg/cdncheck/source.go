package cdncheck

import (
	_ "embed"
	"strings"
)

//go:embed data.json
var data string
var generatedData InputCompiled

func mapKeys(m map[string][]string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return strings.Join(keys, ", ")
}
