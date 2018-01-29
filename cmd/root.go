package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "medkit",
	Short: "MEDKIT is a multi-environment dotfiles manager",
	Long:  `MEDKIT (Multi-Environment Dotfiles Kit) is the dotfile management solution for the obsessive compulsive.`,
	Version: "0.0.1-alpha",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("rootCmd dotfilesDirectory from Run = " + viper.GetString("dotfilesDirectory"))
    },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.medkit.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

    fmt.Println("rootCmd dotfilesDirectory from init = " + viper.GetString("dotfilesDirectory"))
    viper.SetDefault("dotfilesDirectory", "~/dotfiles")
    rootCmd.PersistentFlags().String("dotfilesDirectory", viper.GetString("dotfilesDirectory"), "Path the your dotfiles directory (default is $HOME/dotfiles)")
    viper.BindPFlag("dotfilesDirectory", rootCmd.PersistentFlags().Lookup("dotfilesDirectory"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
    fmt.Println("initiConfig STARTED")
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".medkit" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".medkit")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
    fmt.Println("initiConfig FINISHED")
}
