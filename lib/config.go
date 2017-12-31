package lib

import (
	"os/user"
	"log"

	"github.com/zpatrick/go-config"
)

func InitConfig() *config.Config {
	usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
	}

	yamlFile := config.NewYAMLFile( usr.HomeDir + "/.leetcli.yaml")
    return config.NewConfig([]config.Provider{yamlFile})
}
