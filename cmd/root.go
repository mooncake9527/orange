package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/mooncake9527/orange/cmd/gen"
	"github.com/mooncake9527/orange/cmd/start"
	"github.com/mooncake9527/orange/cmd/version"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:          "orange",
		Short:        "orange",
		Long:         `orange`,
		SilenceUsage: true,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				tip()
				return errors.New("one arg required at least")
			}
			return nil
		},
		PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
		Run: func(cmd *cobra.Command, args []string) {
			tip()
		},
	}
)

func tip() {
	welcomeNotice := `Hi body, welcome to use orange`
	fmt.Printf("%s\n", welcomeNotice)
}

func init() {
	rootCmd.AddCommand(start.Cmd)
	rootCmd.AddCommand(gen.Cmd)
	rootCmd.AddCommand(version.Cmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
