package storage

import (
	"auto-clicker/model"
	"encoding/json"
	"io/ioutil"
)

const (
	path        = "config.raw.json"
	permissions = 0644
)

type File struct{}

func (f File) Save(config model.Config) error {
	file, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, file, permissions)
	if err != nil {
		return err
	}

	return nil
}
