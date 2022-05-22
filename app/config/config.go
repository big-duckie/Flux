package config

import (
	"crypto"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"os"
	"path/filepath"

	"github.com/mcuadros/go-defaults"
	"github.com/pelletier/go-toml/v2"
)

type Golder struct {
	TEMP string
}

type Database struct {
	DB_DRIVER string `default:"sqlite" comment:"Which driver to use for the database.\n# \"sqlite\", \"mysql\", \"postgres\""`
	DB_URI    string `default:"golder.db" comment:"URI to connect to the database."`
}

type Authentication struct {
	SECRET_KEY string `comment:"Do not change this key once set."`
}

type Configuration struct {
	Golder         Golder
	Database       Database
	Authentication Authentication
}

func ReadConfig(path string, config *Configuration) error {
	if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
		fileBytes, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		err = toml.Unmarshal(fileBytes, &config)
		if err != nil {
			return err
		}
	}

	defaults.SetDefaults(config)

	if config.Authentication.SECRET_KEY == "" {
		randomInt, _ := rand.Prime(rand.Reader, 64)
		randomHash := crypto.SHA256.New().Sum(randomInt.Bytes())
		config.Authentication.SECRET_KEY = hex.EncodeToString(randomHash)

		err := SaveConfig(path, config)
		if err != nil {
			return err
		}
	}

	return nil
}

func SaveConfig(path string, config *Configuration) error {
	fileBytes, err := toml.Marshal(&config)
	if err != nil {
		return err
	}
	os.MkdirAll(filepath.Dir(path), os.ModePerm)
	err = os.WriteFile(path, fileBytes, os.ModePerm)
	return err
}
