/*
Copyright © 2023 Xiaoning Sun 1530358579@qq.com
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sunxiaoning/restful-provider/mock/app"
	"github.com/sunxiaoning/restful-provider/mock/config"
	"github.com/sunxiaoning/restful-provider/pkg/timeutils"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	cfg *config.Config
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:               "mock",
	Short:             "mock server",
	Long:              `mock server make a fake server to mock the result of provider service`,
	DisableAutoGenTag: true,
	SilenceUsage:      true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cfg.Validate(); err != nil {
			return err
		}
		return runMock()
	},
}

func runMock() error {
	app, err := app.NewSnapshot(cfg)
	if err != nil {
		return err
	}
	setupQuitSignalHandler(func() {
		app.Stop()
	})
	return app.Serve()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, initLogger)
	rootCmd.PersistentFlags().StringP("config", "c", "", "path to config file")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
}

func initLogger() {
	level, err := logrus.ParseLevel(cfg.Logger.Level)
	if err != nil {
		panic(err)
	}
	logrus.SetLevel(level)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: timeutils.DefaultTimeFormat,
	})
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	cfg = &config.Config{}
	cfgFile := viper.GetString("config")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("mock")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".mock")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Custom using config file:" + viper.ConfigFileUsed())
	} else {
		panic(err)
	}
	if err := viper.Unmarshal(cfg); err != nil {
		panic("Unmarshal cfg err!")
	}
	fmt.Println("log level: " + cfg.Logger.Level)
}

func setupQuitSignalHandler(handler func()) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		var done bool
		for {
			select {
			case sig := <-signals:
				logrus.Warnf("receive signal: %v", sig)
				if !done {
					done = true
					handler()
					logrus.Warnf("handle signal: %v finish", sig)
				}
			}
		}
	}()
}
