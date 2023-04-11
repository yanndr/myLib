package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path"
)

var (
	cfgFile string
	Version string
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.CompletionOptions = cobra.CompletionOptions{DisableDefaultCmd: true}
	rootCmd.SilenceUsage = true
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.myLib/config.yaml)")
	rootCmd.PersistentFlags().StringP("url", "", "http://localhost:8080/v1", "url of the rest api")
	viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))

	rootCmd.AddCommand(versionCmd)
}

var rootCmd = &cobra.Command{
	Use:   "myLib",
	Short: "myLib CLI is a command line interface program that allows you to manage a personal book collection",
	Long: `myLib CLI is a command line interface program that allows you to manage a personal book collection. It allows you to:
	• Add and manage books into the system, including some basic information about those books (title, author, published date, edition, description, genre)
	• Create and manage collections of books
	• List all books, all collections and filter book lists by author, genre or a range of publication dates.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	Version = version
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	err = viper.WriteConfig()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		cfgPath := path.Join(home, ".myLib")
		viper.AddConfigPath(cfgPath)
		cfgFile = path.Join(cfgPath, "config.yaml")
		viper.SetConfigFile(cfgFile)
	}

	if _, err := os.Stat(cfgFile); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := viper.WriteConfig(); err != nil {
				cobra.CheckErr(err)
			}
		} else {
			cobra.CheckErr(err)
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
