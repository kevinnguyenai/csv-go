/*
Copyright © 2022 Kevin Nguyen <kevin.nguyen.eng@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var date string
var crypto []string

// porfolioCmd represents the porfolio command
var porfolioCmd = &cobra.Command{
	Use:   "porfolio",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("porfolio called")
		if len(date) == 0 {
			fmt.Printf("date:%s\n", viper.GetViper().GetString("date"))
		} else {
			_date, _err := time.Parse(time.RFC3339, date)
			if _err == nil {
				fmt.Printf("date: %d\n", _date.Unix())
				pkg.csvparser.create("transactions.csv")
			}
		}
		if len(crypto) > 0 {
			fmt.Printf("crypto: %s\n", crypto)
		}
	},
}

func init() {
	rootCmd.AddCommand(porfolioCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// porfolioCmd.PersistentFlags().String("foo", "", "A help for foo")
	// example crypto : "BTC","ETH"
	porfolioCmd.PersistentFlags().StringArrayVarP(&crypto, "crypto", "c", []string{}, "input crypto which want to sum up")
	// example "2022-Aug-25"
	porfolioCmd.PersistentFlags().StringVarP(&date, "date", "t", "", "input date which want to sum up")
	viper.SetDefault("date", time.Now().Unix())

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// porfolioCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
