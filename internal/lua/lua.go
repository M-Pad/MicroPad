package lua

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/kirsle/configdir"
	lua "github.com/yuin/gopher-lua"
)

var l *lua.LState

// Loads a plugin by name from micropad plugin directory
func LoadPlugin(name string) (*PluginMeta, error) {
	paths := []string{
		filepath.Join(name, "plugin.lua"),                                               // Current directory; usefull for plugin dev
		filepath.Join(configdir.LocalConfig("micropad"), "plugins", name, "plugin.lua"), // Usual plugin location
	}

	var pluginPath string
	for _, path := range paths {
		_, err := os.Stat(path)
		if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}

		if err != nil {
			pluginPath = path
			break
		}
	}

	if pluginPath == "" {
		return nil, os.ErrNotExist
	}

	return nil, nil
}
