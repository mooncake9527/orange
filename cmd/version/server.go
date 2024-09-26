package version

import (
	"fmt"
	"github.com/mooncake9527/orange/common/consts"

	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:     "version",
		Short:   "version",
		Example: "orange version",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	fmt.Println(consts.Version)
	return nil
}
