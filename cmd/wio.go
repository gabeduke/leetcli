package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"leetcli/lib"
	"leetcli/lib/wio/sensors"

	"github.com/spf13/cobra"
)

var wio string

var wioCmd = &cobra.Command{
	Use:   "wio [Wio Name]",
	Short: "get sensor value(s) for [Wio Name]",
	Long: `Pass the Wio name from your config file to get back any sensor
data attached to the node`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := lib.InitConfig()

		sensor, _ := conf.String(wio +".sensor")
		params := sensors.SensorMap(sensor)

		baseurl, _ := conf.String("baseurl")
		token, _ := conf.String(wio + ".token")
		node, _ := conf.String( wio + ".sensor")
		url := (baseurl + node + params + "?access_token=" + token)

		fmt.Printf(callWio(url))
	},
}

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

func init() {
	RootCmd.AddCommand(wioCmd)
	wioCmd.PersistentFlags().StringVarP(&wio, "name", "n", "myWio", "name of wio board")
}
