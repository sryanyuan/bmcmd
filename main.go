package main

import (
	"github.com/spf13/cobra"
	"github.com/sryanyuan/bmcmd/cmd"
)

func main() {
	var rootCmd = &cobra.Command{Use: "bmcmd"}
	rootCmd.AddCommand(cmd.CreateCmdTolua())
	rootCmd.AddCommand(cmd.CreateCmdMergeINI())
	rootCmd.Execute()
}
