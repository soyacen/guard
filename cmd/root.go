package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/yacen/guard/config"
	"github.com/yacen/guard/db"
	"github.com/yacen/guard/server"
	"github.com/yacen/guard/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "guard",
	Short: "Account server based on go and JWT development",
	Long:  `Account server based on go and JWT development.`,
	Run: func(cmd *cobra.Command, args []string) {
		db.InitMysql()
		db.InitRedis()
		util.InitJwtKeyFile()
		server.Start()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// 配置文件目录
	rootCmd.PersistentFlags().StringVar(&config.Cfg.CfgFile, "config", "", "config file (default is $PWD/guard.yaml)")

}

func initConfig() {
	if config.Cfg.CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(config.Cfg.CfgFile)
	} else {
		// Find work directory.
		wd, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(wd)
		// Search config in home directory with name ".guard" (without extension).
		viper.AddConfigPath(wd)
		viper.SetConfigName("guard")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	config.GetFlagsFromConfigFile()
}
