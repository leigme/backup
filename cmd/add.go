/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("add file or dir is nil")
			return
		}
		for _, f := range args {
			if fs, err := os.Stat(f); err == nil {
				if fs.IsDir() {
					saveDir(f)
				} else {
					saveFile(f)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

}

func saveDir(dir string) {}

func saveFile(file string) {
	//files := viper.GetStringSlice(FILES)
}

func contains(s string, arr []string) bool {
	for _, ar := range arr {
		if strings.EqualFold(s, ar) {
			return true
		}
	}
	return false
}
