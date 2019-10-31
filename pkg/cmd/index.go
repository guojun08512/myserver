package cmd

import (
	"errors"

	"myserver/pkg/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ErrUsage is returned by the cmd.Usage() method
var ErrUsage = errors.New("Bad usage of command")
var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "server",
	Short: "server is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
				  love by spf13 and friends in Go.
				  Complete documentation is available at http://hugo.spf13.com`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return config.Setup(cfgFile)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// Do Stuff Here
		return cmd.Usage()
	},
	SilenceUsage: true,
}

func init() {
	usageFunc := RootCmd.UsageFunc()

	RootCmd.SetUsageFunc(func(cmd *cobra.Command) error {
		usageFunc(cmd)
		return ErrUsage
	})
	flags := RootCmd.PersistentFlags()
	flags.StringVarP(&cfgFile, "config", "c", "", "configuration file (default \"$HOME/my-config.yaml\")")
	// flags.String("host", "localhost", "server host")
	// checkNoErr(viper.BindPFlag("host", flags.Lookup("host")))

	flags.IntP("port", "p", 8080, "server port")
	checkNoErr(viper.BindPFlag("port", flags.Lookup("port")))
}

func checkNoErr(err error) {
	if err != nil {
		panic(err)
	}
}
