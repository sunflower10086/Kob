package cmd

import (
	"errors"
	"fmt"
	"snake/version"

	"github.com/spf13/cobra"
)

var (
	vers bool
)

var RootCmd = &cobra.Command{
	Use:     "code",
	Long:    "code running 微服务",
	Short:   "code running 微服务",
	Example: "code commands",
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
