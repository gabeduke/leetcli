package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"leetcli/lib"

	"github.com/spf13/cobra"
)

var wioCmd = &cobra.Command{
	Use:   "wio [Wio Name]",
	Short: "get sensor value(s) for [Wio Name]",
	Long: `Pass the Wio name from your config file to get back any sensor
data attached to the node`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := lib.InitConfig()

		wio := args[0]
		baseurl, _ := conf.String("baseurl")
		token, _ := conf.String(wio + ".token")
		node, _ := conf.String( wio + ".sensor")
		url := (baseurl + node +"/" + "lux" + "?access_token=" + token)

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
