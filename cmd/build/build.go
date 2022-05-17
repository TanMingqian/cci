/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package build

import (
	"fmt"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("build called")
	},
	Args: cobra.ExactArgs(4), // cci build [system] [branch] [workNo] [path]
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("pre run")
	},
}

func NewBuildCommand() *cobra.Command {
	return buildCmd
}

func runBuild(args []string) {

}
