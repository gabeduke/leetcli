package lib

import (
	"os/user"
	"log"

	"github.com/zpatrick/go-config"
)

func InitConfig() *config.Config {
<<<<<<< HEAD
	yamlFile := config.NewYAMLFile("~/.leetcli.yml")
=======
	usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
	}

	yamlFile := config.NewYAMLFile( usr.HomeDir + "/.leetcli.yaml")
>>>>>>> 163b42618b3005510bf749bb1b0c9f56c5601d94
    return config.NewConfig([]config.Provider{yamlFile})
}
