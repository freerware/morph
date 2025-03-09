package morph

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// Loaders is the collection of loaders by file extension.
var Loaders = map[string]Loader{
	"yaml": YAMLLoader{},
	"yml":  YAMLLoader{},
	"json": JSONLoader{},
}

// Load loads the configuration from the provided file.
func Load(path string) (Configuration, error) {
	extension := path[strings.LastIndex(path, ".")+1:]
	loader, ok := Loaders[extension]
	if !ok {
		return Configuration{}, fmt.Errorf("no loader for files with %q extension", extension)
	}
	return loader.Load(path)
}

// Loader loads the cofiguration from the provided file.
type Loader interface {
	Load(path string) (Configuration, error)
}

type JSONLoader struct{}

// Load loads the configuration from the JSON file provided.
func (l JSONLoader) Load(path string) (c Configuration, err error) {
	var file []byte
	if file, err = os.ReadFile(path); err != nil {
		return
	}
	return c, json.Unmarshal(file, &c)
}

type YAMLLoader struct{}

// Load loads the configuration from the YAML file provided.
func (l YAMLLoader) Load(path string) (c Configuration, err error) {
	var file []byte
	if file, err = os.ReadFile(path); err != nil {
		return
	}
	return c, yaml.Unmarshal(file, &c)
}
