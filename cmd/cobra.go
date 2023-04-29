package cmd

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"go-admin/cmd/app"
	"os"

	"github.com/spf13/cobra"

	"go-admin/cmd/api"
	"go-admin/cmd/config"
	"go-admin/cmd/migrate"
	"go-admin/cmd/version"
)

var rootCmd = &cobra.Command{
	Use:          "go-admin",
	Short:        "go-admin",
	SilenceUsage: true,
	Long:         `go-admin`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New(pkg.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
	rootCmd.AddCommand(migrate.StartCmd)
	rootCmd.AddCommand(version.StartCmd)
	rootCmd.AddCommand(config.StartCmd)
	rootCmd.AddCommand(app.StartCmd)
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
