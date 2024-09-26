package gen

import (
	"fmt"
	"github.com/mooncake9527/orange/cmd/com"
	cons "github.com/mooncake9527/orange/common/consts"
	"github.com/mooncake9527/orange/modules/tools/apis"
	"github.com/mooncake9527/orange/modules/tools/service"

	"github.com/spf13/cobra"
	_ "github.com/spf13/viper/remote"
)

var (
	configYml   string
	dbName      string
	tableName   string
	packageName string
	force       bool
	Cmd         = &cobra.Command{
		Use:     "gen",
		Short:   "generate code",
		Long:    "generate code based on database tables",
		Example: "orange gen -c resources/config.dev.yml -d sys -t sys_users",
		Run: func(cmd *cobra.Command, args []string) {
			gen()
		},
	}
)

func init() {
	Cmd.PersistentFlags().StringVarP(&configYml, "config", "c", "resources/config.dev.yml", "Start server with provided configuration file")
	Cmd.PersistentFlags().StringVarP(&dbName, "db", "d", "default", "database name")
	Cmd.PersistentFlags().StringVarP(&tableName, "table", "t", "", "table name")
	Cmd.PersistentFlags().StringVarP(&packageName, "module", "m", "", "module name")
	Cmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "if set to true, will overwrite existing files")
}

func gen() {
	com.Pre(configYml)
	fmt.Printf("modulename %s db %s table %s\n", packageName, dbName, tableName)
	tab, err := service.SerGenTables.GenTableInit(dbName, tableName, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	if packageName != "" {
		tab.PackageName = packageName
	}
	tab.ApiRoot = cons.ApiRoot

	for i, v := range tab.Columns {
		tab.Columns[i].TsType = apis.TypeGo2Ts(v.GoType)
	}
	err = service.SerGenTables.NOMethodsGen(tab, force)
	if err != nil {
		fmt.Println(err)
		return
	}
}
