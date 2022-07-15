/*
Copyright © 2022 Laurean Ray Bahala <laureanraybahala@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// searchCmd represents the search command
var (
	cfgFile string

	searchCmd = &cobra.Command{
		Use:   "search",
		Short: "Search Project/Repos",
		Long: `IPNA is a TUI tool to search and check if a project name is already been used
across github, npm, and other package hosting services`,
		Run: func(cmd *cobra.Command, args []string) {
			Search(cmd, args)
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.AddCommand(searchCmd)
}

func Search(cmd *cobra.Command, args []string) {
	log.Println(args)
	// api.SearchGithubRepos([2])
}

func initConfig() {
	log.Println("Init called!")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {

		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.SetDefault("github_token", "<YOUR_GITHUB_TOKEN_HERE>")
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("ipna")
		viper.SafeWriteConfigAs(home + "/ipna.yaml")
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file: ", viper.ConfigFileUsed())
		fmt.Println("Using token: ", viper.Get("github_token"))
	} else {
		log.Print(err)
	}
}
