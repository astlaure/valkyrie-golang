package core

import (
	"encoding/json"
	"io"

	"github.com/astlaure/valkyrie-golang/web"
)

type (
	ViteAssets struct {
		Links   []string
		Scripts []string
	}
)

func GetViteAssets() ViteAssets {
	manifest, err := web.StaticFS.Open("static/dist/.vite/manifest.json")

	if err != nil {
		panic("Cannot read the manifest")
	}

	defer manifest.Close()

	bytes, _ := io.ReadAll(manifest)

	var assets = ViteAssets{Links: []string{}, Scripts: []string{}}
	var result map[string]interface{}
	json.Unmarshal(bytes, &result)

	keys := make([]string, 0, len(result))
	for k := range result {
		keys = append(keys, k)
	}

	for _, key := range keys {
		element := result[key].(map[string]interface{})
		assets.Scripts = append(assets.Scripts, element["file"].(string))

		if element["css"] != nil {
			for _, item := range element["css"].([]interface{}) {
				assets.Links = append(assets.Links, item.(string))
			}
		}
	}

	return assets
}
