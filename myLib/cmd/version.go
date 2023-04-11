package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"myLib/client"
	"net/http"
	"os"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the version of the program",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cli := client.NewClient(viper.GetString("url"), &http.Client{})
		apiVersion, err := cli.GetVersion(context.TODO())
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		fmt.Printf("myLib: %s api: %s \n", Version, apiVersion)
	},
}
