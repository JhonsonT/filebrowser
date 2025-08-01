package cmd

import (
	"github.com/spf13/cobra"

	"github.com/filebrowser/filebrowser/v2/storage/bolt/importer"
)

func init() {
	rootCmd.AddCommand(upgradeCmd)

	upgradeCmd.Flags().String("old.database", "", "")
	upgradeCmd.Flags().String("old.config", "", "")
	_ = upgradeCmd.MarkFlagRequired("old.database")
}

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrades an old configuration",
	Long: `Upgrades an old configuration. This command DOES NOT
import share links because they are incompatible with
this version.`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		flags := cmd.Flags()
		oldDB, err := getString(flags, "old.database")
		if err != nil {
			return err
		}
		oldConf, err := getString(flags, "old.config")
		if err != nil {
			return err
		}
		db, err := getString(flags, "database")
		if err != nil {
			return err
		}
		return importer.Import(oldDB, oldConf, db)
	},
}
