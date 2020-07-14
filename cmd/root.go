// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package cmd

import (
	"github.com/kunnos/homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/klog"
	"strings"
)

var (
	cfgFile string

	kcCmd = &cobra.Command{
		Use:   "kc",
		Short: "k8s check tools",
		Run: func(cmd *cobra.Command, args []string) {
			klog.Info("start node check")
		},
	}
)

func Execute() error {
	return kcCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	kcCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kc.yaml)")
	kcCmd.AddCommand(versionCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		home, _ := homedir.Dirv2(cfgFile)
		viper.AddConfigPath(home)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			panic(err)
		}
		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".kc")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("kc")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err == nil {
		klog.Info("Using config file:", viper.ConfigFileUsed())
	}

}
