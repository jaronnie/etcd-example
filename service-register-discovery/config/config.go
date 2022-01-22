package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/jaronnie/mysrd/util"
	"github.com/spf13/viper"
)

const (
	defaultReadMode      = "local"
	defaultFilePath      = "./cfg.toml"
	defaultServiceName   = "other"
	defaultEtcdKeyFormat = "/%s/config/cfg.toml"
	defaultEtcdHost      = "127.0.0.1:2379"
)

// CfgManager define config manager
type CfgManager struct {
	readMode string // viper read mode e.g. etcd or local
	host     string // etcd host e.g. "127.0.0.1:2379"
	filePath string // etcd configure key e.g. "/b20/config/cfg.toml"
}

// NewConfigManager new config manager
func NewConfigManager(mode, host, name string) *CfgManager {
	if mode == "" {
		mode = defaultReadMode
	}
	if host == "" {
		host = defaultEtcdHost
	}
	if name == "" {
		name = defaultServiceName
	}

	filePath := fmt.Sprintf(defaultEtcdKeyFormat, name)

	cm := &CfgManager{
		readMode: mode,
		host:     host,
		filePath: filePath,
	}
	return cm
}

// NewConfig new Configure
func (cm *CfgManager) NewConfig() error {
	return cm.newLocalConfig()
}

// newLocalConfig return local configure instance
func (cm *CfgManager) newLocalConfig() (err error) {

	// set local file path
	viper.SetConfigFile(defaultFilePath)
	ext, err := util.Ext(defaultFilePath, viper.SupportedExts...)
	if err != nil {
		panic(err)
	}
	viper.SetConfigType(ext)

	if err = viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// watch config and set onchange function
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// TODO: local file change
		fmt.Println(fmt.Sprintf("Config file changed: %s\n", e.Name))
	})
	return
}
