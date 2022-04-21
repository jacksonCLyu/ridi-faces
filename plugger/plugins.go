package plugger

import (
	"errors"
	"plugin"
)

// Plugin is a plugin loaded from a file.
type Plugin struct {
	// Name Plugin name
	Name string
	// Type Plugin type
	Type string
	// Path Plugin path
	Path string
	// NewFunc Plugin new instance func
	NewFunc interface{}
}

// Load loading plugin build by `go build --buildmode=plugin` from filepath
//return plugin instance and error
func Load(path string) (*Plugin, error) {
	plug, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}
	symbol, err := plug.Lookup("Plugin")
	if err != nil {
		return nil, err
	}
	comp, ok := symbol.(*Plugin)
	if !ok {
		return nil, errors.New("symbol could not cast Plugin")
	}
	return comp, nil
}
