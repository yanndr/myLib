package cmd

import (
	"api/api"
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"myLib/client"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	lastname, firstname, middleName                      string
	lastnameChanged, firstnameChanged, middleNameChanged bool
)

func init() {
	authorCmd.AddCommand(authorCreateCmd)
	authorCmd.AddCommand(authorDeleteCmd)
	authorCmd.AddCommand(authorUpdateCmd)

	authorCreateCmd.Flags().StringVarP(&firstname, "firstname", "f", "", "Set the firstname (or initials) of the author")
	authorCreateCmd.Flags().StringVarP(&middleName, "middlename", "m", "", "Set the middlename of the author")

	authorUpdateCmd.Flags().StringVarP(&lastname, "lastname", "l", "", "Set the last name of the author")
	authorUpdateCmd.Flags().StringVarP(&firstname, "firstname", "f", "", "Set the first name (or initials) of the author")
	authorUpdateCmd.Flags().StringVarP(&middleName, "middlename", "m", "", "Set the middle name of the author")

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

var authorDeleteCmd = &cobra.Command{
	Use:   "delete <lastname>",
	Short: "Delete an author",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			_ = cmd.Usage()
			return nil
		}
		cli := client.NewClient(viper.GetString("url"), &http.Client{})

		lastname := args[0]
		deleted, err := executeActionForAuthors(cli, lastname, "delete", func(author api.Author) error {
			err := cli.DeleteAuthor(context.TODO(), author.ID)
			if err != nil {
				return err
			}
			return nil
		})

		if err != nil {
			return err
		}
		if deleted {
			fmt.Printf("Author %s deleted.\n", lastname)
		}

		return nil
	},
}

var authorUpdateCmd = &cobra.Command{
	Use:   "update <lastname>",
	Short: "Update an author",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			_ = cmd.Usage()
			return nil
		}
		cli := client.NewClient(viper.GetString("url"), &http.Client{})

		lastnameArg := args[0]

		updated, err := executeActionForAuthors(cli, lastnameArg, "use", func(author api.Author) error {
			err := cli.UpdateAuthor(context.TODO(), author.ID, lastname, firstname, middleName, lastnameChanged, firstnameChanged, middleNameChanged)
			if err != nil {
				return err
			}
			return nil
		})

		if err != nil {
			if errors.Is(err, client.DuplicateResourceErr) {
				name := strings.Join(strings.Fields(fmt.Sprintf("%v %v %v", firstname, middleName, lastnameArg)), " ")
				fmt.Printf("Cannot update %s because an Author with the same first name, middle name and last name already exists.\n", name)
				return nil
			}
			return err
		}

		if lastnameChanged {
			lastnameArg = lastname
		}
		if updated {
			fmt.Printf("Author %s updated.\n", lastnameArg)
		}

		return nil
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		lastnameChanged = cmd.Flags().Changed("lastname")
		firstnameChanged = cmd.Flags().Changed("firstname")
		middleNameChanged = cmd.Flags().Changed("middlename")

		if !lastnameChanged && !firstnameChanged && !middleNameChanged {
			return errors.New("at least one flag must be provided")
		}

		return nil
	},
}

func executeActionForAuthors(cli *client.Client, requestName, verb string, fn func(author api.Author) error) (bool, error) {
	authors, err := cli.GetAuthors(context.TODO(), requestName)
	if err != nil {
		return false, err
	}
	l := len(authors)
	if l == 0 {
		fmt.Printf("No authors with lastname %s found\n", requestName)
		return false, nil
	} else if l == 1 {
		err = fn(authors[0])
		if err != nil {
			return false, nil
		}
	} else {
		val, err := multiAuthorsPrompt(requestName, authors, verb)
		if err != nil {
			return false, err
		}
		if val > 0 {
			err := fn(authors[val-1])
			if err != nil {
				return false, err
			}
		}
	}

	return true, nil
}

func multiAuthorsPrompt(lastname string, authors []api.Author, verb string) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%v authors exist with the last name %s:\n", len(authors), lastname)
	for i, a := range authors {
		name := strings.Join(strings.Fields(fmt.Sprintf("%v %v %v", a.FirstName, a.MiddleName, a.LastName)), " ")
		fmt.Printf("[%v] %v\t\n", i+1, name)
	}
	valid := false
	var choice int
	for !valid {
		fmt.Printf("Choose which one you want to %s or enter c to cancel.", verb)
		str, _, err := reader.ReadLine()

		if err != nil {
			return -1, err
		}
		if strings.ToLower(string(str)) == "c" {
			return -1, nil
		}
		choice, err = strconv.Atoi(string(str))
		if err == nil {
			valid = choice <= len(authors) && choice > 0
		}
	}
	return choice, nil
}
