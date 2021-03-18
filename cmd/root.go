package cmd

import (
	"fmt"
	"log"
	"vvvstore/internal/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"vvvstore/internal/pkg/config"
)

var (
	conf string
	host string
	port int
	cors bool
	mode string
	rootCmd = &cobra.Command{
		Use: "vvvstore",
		Short: "vvvstore API server",
		Long: fmt.Sprintf("%s\nStart vvvstore API server", welcome()),
		Run: func(cmd *cobra.Command, args []string) {
			if mode != "debug" && mode != "release" {
				cmd.Help()
				return
			}
			fmt.Println(welcome())
			app.Initialize()
			app.Run()
		},
	}
)


func init()  {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVarP(&conf, "conf", "c", "conf/config.toml", "start server with a configuration file")
	rootCmd.Flags().StringVarP(&mode, "mode", "m", "release", "app running mode [debug|release]")
	rootCmd.Flags().StringVarP(&host, "host", "", "", "listen on tcp host")
	rootCmd.Flags().IntVarP(&port, "port", "p", 9595, "listen on tcp port")
	rootCmd.Flags().BoolVarP(&cors, "cors", "x", false, "enable cross-domain requests")

	viper.BindPFlag("app.port", rootCmd.Flags().Lookup("port"))
	viper.BindPFlag("app.host", rootCmd.Flags().Lookup("host"))
	viper.BindPFlag("app.mode", rootCmd.Flags().Lookup("mode"))
	viper.BindPFlag("app.cors", rootCmd.Flags().Lookup("cors"))
}

// 配置初始化
func initConfig()  {
	if err := config.InitConfig(conf); err != nil {
		log.Fatal(err)
	}
}

// 欢迎
func welcome() string {
	welcomeStr := `
                          _                 
                         | |                
 __   ____   ____   _____| |_ ___  _ __ ___ 
 \ \ / /\ \ / /\ \ / / __| __/ _ \| '__/ _ \
  \ V /  \ V /  \ V /\__ \ || (_) | | |  __/
   \_/    \_/    \_/ |___/\__\___/|_|  \___|
                                             
`
	return welcomeStr
}


func Execute() error {
	return rootCmd.Execute()
}