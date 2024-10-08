package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
)

type Exchange struct {
	Data []struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"data"`
}

// exchangeCmd represents the exchange command
var exchangeCmd = &cobra.Command{
	Use:   "exchange",
	Short: "A brief description of your command",
	Long:  ``,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fundType := args[0]

		url := fmt.Sprintf("https://api.twelvedata.com/exchanges?type=%s", fundType)
		req, err := http.NewRequest("GET", url, nil)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			panic("Stock API not Available")
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		var exchangeResponse Exchange
		err = json.Unmarshal(body, &exchangeResponse)
		if err != nil {
			panic(err)
		}

		for _, exchange := range exchangeResponse.Data {
			fmt.Printf("Name: %s, Country: %s\n", exchange.Name, exchange.Country)
		}
	},
}

func init() {
	rootCmd.AddCommand(exchangeCmd)
}
