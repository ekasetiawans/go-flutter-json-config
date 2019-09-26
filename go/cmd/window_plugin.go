package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const channelName = "com.github.ekasetiawans/window_plugin"

type windowPlugin struct{}

var _ flutter.Plugin = &windowPlugin{} 
var _ flutter.PluginGLFW = &windowPlugin{} 

func (p *windowPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})

	channel.HandleFunc("setFullScreen", handleSetFullScreen)
	channel.HandleFunc("getFullScreen", handleGetFullScreen)

	return nil
}

func handleGetFullScreen(arguments interface{}) (reply interface{}, err error) {
	config := getConfiguration()
	return config.FullScreenMode, nil
}

func handleSetFullScreen(arguments interface{}) (reply interface{}, err error) {
	value := arguments.(bool)
	config := getConfiguration()
	config.FullScreenMode = value
	config.saveConfig()

	return true, nil
}

// InitPluginGLFW .
func (p *windowPlugin) InitPluginGLFW(window *glfw.Window) error {
	return nil
}

type jsonConfiguration struct {
	FullScreenMode bool    `json:"fullscreen"`
}

func getConfiguration() *jsonConfiguration {
	filename := "config.json"
	if _, err := os.Stat(filename); err == nil {
		b, err := ioutil.ReadFile(filename)
		if err == nil {
			var config jsonConfiguration
			json.Unmarshal(b, &config)
			return &config
		}
	}

	return &jsonConfiguration{
		FullScreenMode: false,
	}
}

func (config *jsonConfiguration) saveConfig() {
	b, err := json.Marshal(config)
	if err != nil {
		log.Println(err)
		return
	}

	filename := "config.json"
	ioutil.WriteFile(filename, b, 0644)
}


func getWindowModeFromJSON() flutter.Option {
	config := getConfiguration()
	if config.FullScreenMode {
		return flutter.WindowMode(flutter.WindowModeBorderlessFullscreen)
	}

	return flutter.WindowMode(flutter.WindowModeDefault)
}