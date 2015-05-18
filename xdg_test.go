package xdg_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	xdg "."
	"github.com/stretchr/testify/assert"
)

func TestConfigDirs(t *testing.T) {
	os.Clearenv()

	assert.Equal(t, []string{"/etc/xdg"}, xdg.ConfigDirs())

	os.Setenv("XDG_CONFIG_DIRS", "/this/that:some/place:/what")
	assert.Equal(t, []string{"/this/that", "/what"}, xdg.ConfigDirs())
}

func TestConfigHome(t *testing.T) {
	os.Clearenv()

	assert.Equal(t, os.ExpandEnv("$HOME/.config"), xdg.ConfigHome())

	os.Setenv("XDG_CONFIG_HOME", "/some/other/place/.config")
	assert.Equal(t, "/some/other/place/.config", xdg.ConfigHome())
}

func TestConfigWithHome(t *testing.T) {
	os.Clearenv()
	os.Setenv("XDG_CONFIG_HOME", os.TempDir())

	f, _ := ioutil.TempFile("", "")

	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()

	assert.Equal(t, f.Name(), xdg.Config(filepath.Base(f.Name())))
}

func TestConfigWithDirs(t *testing.T) {
	os.Clearenv()
	os.Setenv("XDG_CONFIG_DIRS", os.TempDir())

	f, _ := ioutil.TempFile("", "")

	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()

	assert.Equal(t, f.Name(), xdg.Config(filepath.Base(f.Name())))
}

func TestConfigWhenNotfound(t *testing.T) {
	os.Clearenv()
	os.Setenv("XDG_CONFIG_DIRS", os.TempDir())

	assert.Equal(t, "", xdg.Config("/tmp/this-probably-doesnt-exist"))
}
