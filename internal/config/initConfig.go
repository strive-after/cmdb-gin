package config

import (
	"github.com/BurntSushi/toml"
)

func InitConfig(fpath string) (*Configuration, error) {
	var CmdbConfig = Configuration{}
	_, err := toml.DecodeFile(fpath, &CmdbConfig)
	if err != nil {
		return nil, err
	}
	return &CmdbConfig, nil
}
