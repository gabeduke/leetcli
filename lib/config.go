package lib

import (
	"github.com/zpatrick/go-config"
)

func InitConfig() *config.Config {
	yamlFile := config.NewYAMLFile("~/.leetcli.yml")
    return config.NewConfig([]config.Provider{yamlFile})
}
