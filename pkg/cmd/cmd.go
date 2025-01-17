package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/czar0/fabkit-cli/pkg/cmd/generate"
	"github.com/czar0/fabkit-cli/pkg/cmd/network"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:     "fabkit",
		Aliases: []string{"fk"},
		Short:   "Fabkit CLI is an utility to interact with Hyperledger Fabric",
		Long:    getTitle() + `Fabkit CLI is an utility to interact with Hyperledger Fabric.`,
	}
)

func getTitle() string {
	return `
	╔═╗┌─┐┌┐ ┬┌─ ┬┌┬┐
	╠╣ ├─┤├┴┐├┴┐ │ │ 
	╚  ┴ ┴└─┘┴ ┴ ┴ ┴ 
              ■-■-■               
`
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is <current_dir>/config.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "Cesare Valitutto", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	if err := viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author")); err != nil {
		log.Fatalln(err)
	}
	if err := viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper")); err != nil {
		log.Fatalln(err)
	}
	viper.SetDefault("author", "Cesare Valitutto <cesare.valitutto@gmail.com>")
	viper.SetDefault("license", "apache")

	rootCmd.AddCommand(network.NewCmdNetwork())
	rootCmd.AddCommand(generate.NewGenerateCmd())
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.Getwd()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
