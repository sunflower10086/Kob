package cmd

import (
	"backend/version"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	vers bool
)

var RootCmd = &cobra.Command{
	Use:     "backend",
	Long:    "backend后端",
	Short:   "backend后端",
	Example: "backend后端 commands",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println(version.FullVersion())
			return nil
		}
		return errors.New("no flags find")
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "domo version")
}
