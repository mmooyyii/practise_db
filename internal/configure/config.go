package configure

import (
	"exercise_db/internal/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
)

type Config struct {
	Network network `yaml:"network"`
	Storage storage `yaml:"storage"`
}

func (n *network) ToAddress() string {
	return n.Host + ":" + strconv.Itoa(n.Port)
}

type network struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type storage struct {
	Path string `yaml:"path"`
}

func NewConfig(configPath string) *Config {
	data, err := ioutil.ReadFile(configPath)
	utils.Crash(err)
	config := &Config{}
	err = yaml.Unmarshal(data, config)
	utils.Crash(err)
	return config
}
