package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

var (
	userConf, _ = os.UserConfigDir()
	confPath    = fmt.Sprint(userConf, "/fujira.conf")
)

var conf *Config

type Config struct {
	BasicAuth    *BasicAuth `json:"BasicAuth,omitempty"`
	WorkspaceURL *string
}

func (c *Config) SetBasicAuth(email, token string) {
	ba := &BasicAuth{email, token}
	c.BasicAuth = ba
}

func (c *Config) GetBasicAuth() (BasicAuth, error) {
	if c.BasicAuth == nil {
		return BasicAuth{}, errors.New("Basic auth not configured")
	}

	return *(c.BasicAuth), nil
}

func (c *Config) GetWorkspaceURL() (string, error) {
	if c.WorkspaceURL == nil {
		return "", errors.New("Workspace URL not cofigured")
	}

	return *(c.WorkspaceURL), nil
}

func (c *Config) Save() error {
	buff, err := json.Marshal(c)
	if err != nil {
		log.Println("Error while marshaling config")
		return err
	}

	log.Println("Saving setting: ", string(buff))

	f, err := os.OpenFile(confPath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
	if err != nil {
		log.Println("Cannot open/create conf file: ", confPath)
		return err
	}
	defer f.Close()

	n, err := f.Write(buff)
	if n != len(buff) {
		log.Println("Mismatch len of config and saved data")
		return errors.New("Written len of confing different that buffer len")
	}

	return err
}

type BasicAuth struct {
	Email string
	Token string
}

func (b *BasicAuth) GetEmail() string {
	return b.Email
}

func (b *BasicAuth) GetToken() string {
	return b.Token
}

func loadFromFile() (*Config, error) {
	buff, err := os.ReadFile(confPath)
	if err != nil {
		return nil, err
	}

	var conf Config
	if err := json.Unmarshal(buff, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

func GetConfig() *Config {
	if conf != nil {
		return conf
	}

	if c, err := loadFromFile(); err == nil {
		conf := c
		return conf
	}

	url := "fph.atlassian.net"
	conf = &Config{WorkspaceURL: &url}
	return conf
}
