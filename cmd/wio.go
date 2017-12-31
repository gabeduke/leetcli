package cmd

import (
	"os/user"
	"log"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gabeduke/leetcli-wiolib"
	"github.com/zpatrick/go-config"
	"github.com/spf13/cobra"
)

var wio string

var wioCmd = &cobra.Command{
	Use:   "wio [Wio Name]",
	Short: "get sensor value(s) for [Wio Name]",
	Long: `Pass the Wio name from your config file to get back any sensor
data attached to the node`,
	Run: func(cmd *cobra.Command, args []string) {
		// initialize config
		conf := InitConfig()


		// get sensor param for building url
		sensor, _ := conf.String(wio +".sensor")
		params := sensors.SensorMap(sensor)

		// build url
		baseurl, _ := conf.String("baseurl")
		token, _ := conf.String(wio + ".token")
		node, _ := conf.String( wio + ".sensor")
		url := (baseurl + node + params + "?access_token=" + token)

		// return sensor response
		fmt.Printf(callWio(url))
	},
}

// Execute http request
func callWio(s string) string {

		resp, err := http.Get(s)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		return string(bodyBytes)
}

func InitConfig() *config.Config {
	usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
	}

	yamlFile := config.NewYAMLFile( usr.HomeDir + "/.leetcli.yaml")
    return config.NewConfig([]config.Provider{yamlFile})
}

func init() {
	RootCmd.AddCommand(wioCmd)
	wioCmd.PersistentFlags().StringVarP(&wio, "name", "n", "myWio", "name of wio board")
}
