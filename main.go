package main

import (
	"runtime"

	"github.com/spf13/cobra"

	"github.com/pdxjohnny/s/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var rootCmd = &cobra.Command{Use: "s"}
	rootCmd.AddCommand(commands.Commands...)
	rootCmd.Execute()
}
