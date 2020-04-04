package plugins

import (
	"io/ioutil"
	"os"
	"path"
)

type Handler struct {
	pluginServer     *pluginServer
	pluginsDirectory string
}

func NewHandler() *Handler {
	return &Handler{
		pluginServer: newPluginServer(),
	}
}

func (handler *Handler) Start() error {
	configDirectory, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	handler.pluginsDirectory = path.Join(configDirectory, "nyed/plugins")

	go handler.pluginServer.start()

	return handler.loadPlugins()
}

func (handler *Handler) Stop() error {
	handler.pluginServer.stop()

	return nil
}

func (handler *Handler) loadPlugins() error {
	files, err := ioutil.ReadDir(handler.pluginsDirectory)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			if err := handler.loadPlugin(file.Name()); err != nil {
				return err
			}
		}
	}

	return nil
}

func (handler *Handler) loadPlugin(pluginDirectory string) error {
	pluginConfigPath := path.Join(handler.pluginsDirectory, pluginDirectory, "plugin.toml")

	_ = pluginConfigPath

	return nil
}
