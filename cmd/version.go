package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"runtime"
)

var versionCmd = &cobra.Command{
	Use:"version",
	Short: "print vvvstore version",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		version := viper.GetString("app.version")
		fmt.Println(fmt.Sprintf("vvvstore version v%s %s/%s", version, runtime.GOOS, runtime.GOARCH))
	},
}

func init()  {

	rootCmd.AddCommand(versionCmd)
}