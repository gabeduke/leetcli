package lib

import (
	"github.com/zpatrick/go-config"
)

func InitConfig() *config.Config {
	yamlFile := config.NewYAMLFile("/home/gabeduke/.go/src/leetcli/config.yml")
    return config.NewConfig([]config.Provider{yamlFile})
}
