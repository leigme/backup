/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

const (
	REMOTE = "remote"
	FILES  = "files"
	DIRS   = "dirs"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "backup",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	defer func() {
		saveConfig()
	}()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// 在{home}/config/目录中查找名为"backup.yaml"的配置文件
	viper.AddConfigPath(configDir())
	viper.SetConfigType("yaml")
	viper.SetConfigName("backup")

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}

}

func configDir() string {
	return filepath.Dir(configPath())
}

func configPath() string {
	if h, err := os.UserHomeDir(); err == nil {
		return filepath.Join(h, ".config", "backup.yaml")
	}
	return "backup.yaml"
}

func saveConfig() {
	log.Println("save config")
	if err := viper.WriteConfigAs(configPath()); err != nil {
		log.Println(err)
	}
}
