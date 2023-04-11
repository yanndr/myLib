package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"myLib/client"
	"net/http"
	"strings"
)

var (
	firstname, middleName string
)

func init() {
	authorCmd.AddCommand(authorCreateCmd)

	authorCreateCmd.Flags().StringVarP(&firstname, "firstname", "f", "", "Set the firstname (or initials) of the author")
	authorCreateCmd.Flags().StringVarP(&middleName, "middlename", "m", "", "Set the middlename of the author")

}

var authorCmd = &cobra.Command{
	Use:   "author",
	Short: "Manage authors",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

var authorCreateCmd = &cobra.Command{
	Use:   "create <lastname>",
	Short: "Create an author",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			_ = cmd.Usage()
			return nil
		}
		cli := client.NewClient(viper.GetString("url"), &http.Client{})

		lastnameArg := args[0]
		err := cli.CreateAuthor(context.TODO(), lastnameArg, firstname, middleName)
		if err != nil {
			if errors.Is(err, client.DuplicateResourceErr) {
				name := strings.Join(strings.Fields(fmt.Sprintf("%v %v %v", firstname, middleName, lastnameArg)), " ")
				fmt.Printf("Author %s already exists.\n", name)
				return nil
			}
			return err
		}
		fmt.Printf("Author %s created.\n", lastnameArg)
		return nil
	},
}
