package cdncheck

import (
	_ "embed"
	"encoding/json"
	"fmt"
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

func init() {
	if err := json.Unmarshal([]byte(data), &generatedData); err != nil {
		panic(fmt.Sprintf("Could not parse cidr data: %s", err))
	}
	DefaultCDNProviders = mapKeys(generatedData.CDN)
	DefaultWafProviders = mapKeys(generatedData.WAF)
	DefaultCloudProviders = mapKeys(generatedData.Cloud)
}
