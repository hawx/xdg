// Package xdg provides functions for finding configuration files using the XDG
// Base Directory Standard.
package xdg

import (
	"os"
	"path/filepath"
	"strings"
)

// ConfigDirs returns a ordered list of directories to search for configuration
// files in. It is defined by the environment variable $XDG_CONFIG_DIRS, but if
// this is not present defaults to /etc/xdg.
func ConfigDirs() []string {
	dirs := []string{"/etc/xdg"}

	if env := os.Getenv("XDG_CONFIG_DIRS"); env != "" {
		dirs = dirs[0:0]
		for _, path := range strings.Split(env, ":") {
			if !strings.HasPrefix(path, "/") {
				continue
			}
			dirs = append(dirs, path)
		}
	}

	return dirs
}

// ConfigHome returns the directory configuration files should be stored in,
// this directory has higher precendence than those given by ConfigDirs. It is
// defined by the environment variable $XDG_CONFIG_HOME, but if this is not
// present defaults to $HOME/.config.
func ConfigHome() string {
	home := os.ExpandEnv("$HOME/.config")

	if env := os.Getenv("XDG_CONFIG_HOME"); env != "" {
		home = env
	}

	return home
}

// Config finds the configuration file with the given name. It returns a path
// only if a file exists at that path; if all search directories are exhausted
// then the empty string is returned.
func Config(name string) string {
	p := filepath.Join(ConfigHome(), name)
	if _, err := os.Stat(p); err == nil {
		return p
	}

	for _, dir := range ConfigDirs() {
		p = filepath.Join(dir, name)
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}

	return ""
}
