package storage

import (
	"auto-clicker/model"
	"encoding/json"
	"io/ioutil"
)

const (
	rawPath     = "config.raw.json"
	configPath  = "./config.json"
	permissions = 0644
)

type File struct{}

func (f File) SaveConfig(config model.Config) error {
	file, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(rawPath, file, permissions)
	if err != nil {
		return err
	}

	return nil
}

func (f File) GetConfig() (model.Config, error) {
	plan, err := ioutil.ReadFile(configPath)
	if err != nil {
		return model.Config{}, err
	}

	var conf model.Config
	err = json.Unmarshal(plan, &conf)
	if err != nil {
		return model.Config{}, err
	}

	return conf, nil
}
